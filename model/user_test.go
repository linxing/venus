package model

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_GetByDepartment(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		departmentID := r.Int63n(100) + 1
		user := User{
			UserName:     fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber:  fmt.Sprintf("123456789_%d", r.Intn(100)+1),
			Nickname:     fmt.Sprintf("nickname_%d", r.Int()),
			DepartmentID: departmentID,
		}

		{
			err := UserStatic.Insert(ctx, &user)
			require.NoError(t, err)
		}

		gotUser, err := UserStatic.GetByDepartment(ctx, departmentID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser[0].DeletedAt = time.Time{}
		}

		assert.Equal(t, user, gotUser[0])
	})
}

func TestUser_GetByName(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		gotUser, err := UserStatic.GetByName(ctx, user.UserName)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, user, gotUser)
	})
}

func TestUser_GetByID(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		gotUser, err := UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, user, gotUser)
	})
}

func TestUser_UpdateUserPhoneNumberByID(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}
		newPhoneNumber := "00002222"

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		err = UserStatic.UpdateUserPhoneNumberByID(ctx, user.ID, newPhoneNumber)
		require.NoError(t, err)

		gotUser, err := UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, newPhoneNumber, gotUser.PhoneNumber)
	})
}

func TestUser_UpdateUserRoleByID(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		roleID := r.Int63n(10)
		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		err = UserStatic.UpdateUserRoleByID(ctx, user.ID, roleID)
		require.NoError(t, err)

		gotUser, err := UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, int(roleID), gotUser.RoleID)
	})
}

func TestUser_UpdateUserDepartmentByID(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		departmentID := r.Int63n(100)
		departmentName := fmt.Sprintf("department_%d", r.Int())
		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		err = UserStatic.UpdateUserDepartmentByID(ctx, user.ID, departmentID, departmentName)
		require.NoError(t, err)

		gotUser, err := UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, departmentID, gotUser.DepartmentID)
		assert.Equal(t, departmentName, gotUser.DepartmentName)
	})
}

func TestUser_UpdateUserPositionByID(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should succeed in normal case", func(t *testing.T) {

		db, err := NewDB(User{})
		require.NoError(t, err)

		defer db.Close()
		ctx := context.Background()

		position := fmt.Sprintf("position_%d", r.Int())
		user := &User{
			UserName:    fmt.Sprintf("username_%d", r.Int63()),
			PhoneNumber: fmt.Sprintf("123456789_%d", r.Intn(100)+1),
		}

		{
			err := UserStatic.Insert(ctx, user)
			require.NoError(t, err)
		}

		err = UserStatic.UpdateUserPositionByID(ctx, user.ID, position)
		require.NoError(t, err)

		gotUser, err := UserStatic.GetByID(ctx, user.ID)
		require.NoError(t, err)

		{
			user.CreatedAt = user.CreatedAt.Truncate(time.Second)
			user.UpdatedAt = user.UpdatedAt.Truncate(time.Second)
			gotUser.DeletedAt = time.Time{}
		}

		assert.Equal(t, position, gotUser.Position)
	})
}
