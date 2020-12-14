package user

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"venus/env/envtest"
	"venus/env/global"
	"venus/model"
	"venus/router"
	mockrouter "venus/router/mock"
	mockhttp "venus/testing/http"
	mockrbac "venus/testing/rbac"
	"venus/util"
)

func Test_GetUserInfo(t *testing.T) {

	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	ctx := context.Background()

	t.Run("should success if user exists", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption())
		defer me.Close()

		obj := "/venus/user/info"
		token := ""
		nickname := fmt.Sprintf("nickname_%d", r.Int())
		avatar := fmt.Sprintf("avatar_%d", r.Int())
		position := fmt.Sprintf("position_%d", r.Int())
		departmentName := fmt.Sprintf("department_name_%d", r.Int())
		internalPhone := "10010"

		{
			user := &model.User{
				RoleID:              1,
				Nickname:            nickname,
				Avatar:              avatar,
				Position:            position,
				DepartmentName:      departmentName,
				UserName:            fmt.Sprintf("user_name_%d", r.Int()),
				PhoneNumber:         fmt.Sprintf("phone_number_%d", r.Int()),
				Password:            fmt.Sprintf("password_%d", r.Int()),
				LastLoginAt:         now,
				InternalPhoneNumber: internalPhone,
			}

			err := model.UserStatic.Insert(ctx, user)
			require.NoError(t, err)

			sub := fmt.Sprintf("%d", user.RoleID)
			act := "GET"

			rbac, err := mockrbac.NewRbac(sub, obj, act)
			require.NoError(t, err)

			defer rbac.Close()

			token, err = util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
			require.NoError(t, err)
		}

		rr, err := mockhttp.RequestWithBody(http.MethodGet, obj, nil, token, g)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)

		assert.JSONEq(t, fmt.Sprintf(`
			{
				"avatar": "%s",
				"nickname": "%s",
				"position": "%s",
				"department_name": "%s",
				"last_login_at_sec": %d,
				"internal_phone": "%s",
				"role_id": 1,
				"user_id": 1
			}
		`, avatar, nickname, position, departmentName, now.Unix(), internalPhone), rr.Body.String())
	})
}

func Test_CreateUser(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	phoneNumber := "112233445566"
	userName := fmt.Sprintf("user_name_%d", r.Int())
	password := fmt.Sprintf("password_%d", r.Int())
	nickname := fmt.Sprintf("nickname_%d", r.Int())

	ctx := context.Background()

	t.Run("should success if user not exists", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption())
		defer me.Close()

		jsonBuf, err := json.Marshal(createUserReq{
			UserName:    userName,
			Password:    password,
			PhoneNumber: phoneNumber,
			Nickname:    nickname,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/create", jsonBuf, "", g)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)

		gotUser, err := model.UserStatic.GetByName(ctx, userName)
		require.NoError(t, err)

		assert.Equal(t, phoneNumber, gotUser.PhoneNumber)
	})

	t.Run("should fail if user name exists", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption())
		defer me.Close()

		{
			err := model.UserStatic.Insert(ctx, &model.User{
				UserName:    userName,
				PhoneNumber: fmt.Sprintf("phone_number_%d", r.Int()),
				Password:    fmt.Sprintf("password_%d", r.Int()),
				Nickname:    nickname,
			})
			require.NoError(t, err)
		}

		jsonBuf, err := json.Marshal(createUserReq{
			UserName:    userName,
			Password:    password,
			PhoneNumber: phoneNumber,
			Nickname:    nickname,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/create", jsonBuf, "", g)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if user phone number exists", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption())
		defer me.Close()

		{
			err := model.UserStatic.Insert(ctx, &model.User{
				UserName:    fmt.Sprintf("user_name_%d", r.Int()),
				PhoneNumber: phoneNumber,
				Password:    fmt.Sprintf("password_%d", r.Int()),
				Nickname:    nickname,
			})
			require.NoError(t, err)
		}

		jsonBuf, err := json.Marshal(createUserReq{
			UserName:    fmt.Sprintf("user_name_%d", r.Int()),
			Password:    fmt.Sprintf("password_%d", r.Int()),
			PhoneNumber: phoneNumber,
			Nickname:    nickname,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/create", jsonBuf, "", g)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func Test_UpdateUserPhoneNumberAndNickname(t *testing.T) {

	ctx := context.Background()

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	name := fmt.Sprintf("usernmae_%d", r.Int31())
	newPhoneNumber := fmt.Sprintf("number_%d", r.Int31())
	user := &model.User{
		UserName:    name,
		Nickname:    fmt.Sprintf("nickname_%d", r.Int()),
		Password:    fmt.Sprintf("password_%d", r.Int31()),
		RoleID:      10,
		PhoneNumber: fmt.Sprintf("number_%d", r.Int31()),
	}

	{
		sub := fmt.Sprintf("%d", user.RoleID)
		obj := "/venus/user/info"
		act := "POST"

		rbac, err := mockrbac.NewRbac(sub, obj, act)
		require.NoError(t, err)

		defer rbac.Close()
	}

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	{
		err := model.UserStatic.Insert(ctx, user)
		require.NoError(t, err)
	}

	token, err := util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
	require.NoError(t, err)

	jsonBuf, err := json.Marshal(updatePhoneNumberAndNicknameReq{
		PhoneNumber: newPhoneNumber,
		Nickname:    user.Nickname,
	})
	require.NoError(t, err)

	rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/user/info", jsonBuf, token, g)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, rr.Code)

	gotUser, err := model.UserStatic.GetByID(ctx, user.ID)
	require.NoError(t, err)

	assert.Equal(t, newPhoneNumber, gotUser.PhoneNumber)
}

func Test_SigninUserByName(t *testing.T) {

	ctx := context.Background()

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	userName := fmt.Sprintf("user_name_%d", r.Int31())
	password := fmt.Sprintf("password_%d", r.Int31())

	{
		hashPassword, err := util.HashPassword(password)
		require.NoError(t, err)

		err = model.UserStatic.Insert(ctx, &model.User{
			UserName:    userName,
			Password:    hashPassword,
			RoleID:      10,
			PhoneNumber: "11223344",
		})
		require.NoError(t, err)
	}

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	t.Run("should success in normat case", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByNameReq{
			UserName: userName,
			Password: password,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_name", jsonBuf, "", g)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)

		assert.NotEmpty(t, rr.Body.String())
	})

	t.Run("should fail if wrong password", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByNameReq{
			UserName: userName,
			Password: fmt.Sprintf("wrong_password_%d", r.Int()),
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_name", jsonBuf, "", g)
		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if user not exits", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByNameReq{
			UserName: fmt.Sprintf("not_user_%d", r.Int()),
			Password: password,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_name", jsonBuf, "", g)
		require.NoError(t, err)

		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func Test_SigninUserByPhone(t *testing.T) {

	ctx := context.Background()

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	phoneNumber := "13805034411"
	password := fmt.Sprintf("password_%d", r.Int31())
	{
		hashPassword, err := util.HashPassword(password)
		require.NoError(t, err)

		err = model.UserStatic.Insert(ctx, &model.User{
			UserName:    fmt.Sprintf("user_name_%d", r.Int()),
			Password:    hashPassword,
			RoleID:      r.Intn(100),
			PhoneNumber: phoneNumber,
		})
		require.NoError(t, err)
	}

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	t.Run("should success in normal case", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByPhoneReq{
			PhoneNumber: phoneNumber,
			Password:    password,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_phone", jsonBuf, "", g)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, rr.Code)
		assert.NotEmpty(t, rr.Body.String())
	})

	t.Run("should fall if phone not exists", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByPhoneReq{
			PhoneNumber: fmt.Sprintf("not_phone_%d", r.Int()),
			Password:    password,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_phone", jsonBuf, "", g)
		require.NoError(t, err)

		require.Equal(t, http.StatusNotFound, rr.Code)
		assert.NotEmpty(t, rr.Body.String())
	})

	t.Run("should fall if wrong password", func(t *testing.T) {

		jsonBuf, err := json.Marshal(signinUserByPhoneReq{
			PhoneNumber: phoneNumber,
			Password:    fmt.Sprintf("wrong_password_%d", r.Int()),
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, "/venus/auth/user/signin_with_phone", jsonBuf, "", g)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		assert.NotEmpty(t, rr.Body.String())
	})
}

func Test_GetRefreshToken(t *testing.T) {

	ctx := context.Background()

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))

	user := &model.User{
		UserName:    fmt.Sprintf("usernmae_%d", r.Int31()),
		Nickname:    fmt.Sprintf("nickname_%d", r.Int()),
		Password:    fmt.Sprintf("password_%d", r.Int31()),
		RoleID:      10,
		PhoneNumber: fmt.Sprintf("number_%d", r.Int31()),
	}
	obj := "/venus/user/token/refresh"

	{
		sub := fmt.Sprintf("%d", user.RoleID)
		act := "GET"

		rbac, err := mockrbac.NewRbac(sub, obj, act)
		require.NoError(t, err)

		defer rbac.Close()
	}

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	{
		err := model.UserStatic.Insert(ctx, user)
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	token := ""

	func() {
		monkey.Patch(time.Now, func() time.Time {
			return now.Add(time.Minute)
		})
		defer monkey.UnpatchAll()

		var err error

		token, err = util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
		require.NoError(t, err)
	}()

	{
		httpReq, err := http.NewRequest(http.MethodGet, obj, nil)
		httpReq.Header.Add("x-token", token)

		require.NoError(t, err)

		c.Request = httpReq
	}

	g.HandleContext(c)
	require.Equal(t, http.StatusOK, rr.Code)

	assert.NotEqual(t, fmt.Sprintf(`{"token": "%s"}`, token), rr.Body.String())
}

func Test_EditUserPassword(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	router := router.NewRouter(mockrouter.Servlet{
		User: new(Servlet),
	})
	g := router.GinEngine

	ctx := context.Background()

	t.Run("should success", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption(model.User{}))
		defer me.Close()

		obj := "/venus/user/password"
		token := ""
		user := &model.User{
			UserName:    fmt.Sprintf("name_%d", r.Int()),
			Nickname:    fmt.Sprintf("nickname_%d", r.Int()),
			Password:    fmt.Sprintf("password_%d", r.Int31()),
			RoleID:      10,
			PhoneNumber: fmt.Sprintf("number_%d", r.Int31()),
		}

		password := fmt.Sprintf("password_%d", r.Int())

		{
			err := model.UserStatic.Insert(ctx, user)
			require.NoError(t, err)

			sub := fmt.Sprintf("%d", user.RoleID)
			act := "POST"

			rbac, err := mockrbac.NewRbac(sub, obj, act)
			require.NoError(t, err)

			defer rbac.Close()

			token, err = util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
			require.NoError(t, err)
		}

		jsonBuf, err := json.Marshal(updateUserPasswordReq{
			Password:        password,
			ConfirmPassword: password,
		})
		require.NoError(t, err)

		rr, err := mockhttp.RequestWithBody(http.MethodPost, obj, jsonBuf, token, g)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rr.Code)

		gotUser, err := model.UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		ok := util.CheckPasswordHash(password, gotUser.Password)
		assert.Equal(t, ok, true)
	})
}
