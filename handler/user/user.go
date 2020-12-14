package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"venus/env/global"
	"venus/middleware/auth"
	"venus/model"
	"venus/statusmsg"
	"venus/util"
)

type (
	updatePhoneNumberAndNicknameReq struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		Nickname    string `json:"nickname" binding:"required"`
	}

	createUserReq struct {
		UserName      string `json:"user_name" binding:"required"`
		Nickname      string `json:"nickname" binding:"required"`
		PhoneNumber   string `json:"phone_number" binding:"required"`
		InternalPhone string `json:"internal_phone"`
		Password      string `json:"password" binding:"required"`
	}

	signinUserByNameReq struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	signinUserByPhoneReq struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	signinUserResp struct {
		Token string `json:"token" `
	}

	getUserInfoResp struct {
		Avatar         string `json:"avatar"`
		Nickname       string `json:"nickname"`
		Position       string `json:"position"`
		DepartmentName string `json:"department_name"`
		LastLoginAtSec int64  `json:"last_login_at_sec"`
		InternalPhone  string `json:"internal_phone"`
		RoleID         int    `json:"role_id"`
		UserID         int64  `json:"user_id"`
	}

	updateUserPasswordReq struct {
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)

/**
@api {get} /venus/user/info 获取用户个人信息
@apiGroup User

@apiParam {String} nickname 昵称
@apiParam {String} avatar 头像地址
@apiParam {String} position 职位名称
@apiParam {String} department_name 部门名称
@apiParam {Number} last_login_at_sec 上一次登录时间

@apiSuccessExample Success-Response:
HTTP/1.2 200 OK
{
 "nickname": "test",
 "avatar": "test",
 "position": "test",
 "department_name": "gov",
 "last_login_at_sec": 111112222331,
 "internal_phone": "10020"
 "role_id": 1,
 "user_id": 1
}
*/
func (*Servlet) GetUserInfo(c *gin.Context) {

	userID, _ := auth.GetUserIdAndNickname(c)
	if userID == 0 {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	ctx := c.Request.Context()

	user, err := model.UserStatic.GetByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	resp := getUserInfoResp{
		Avatar:         user.Avatar,
		Nickname:       user.Nickname,
		Position:       user.Position,
		DepartmentName: user.DepartmentName,
		LastLoginAtSec: user.LastLoginAt.Unix(),
		InternalPhone:  user.InternalPhoneNumber,
		RoleID:         user.RoleID,
		UserID:         userID,
	}

	c.JSON(http.StatusOK, resp)
}

/**
@api {post} /venus/auth/user/create 用户注册
@apiGroup User

@apiParam {String} user_name 用户名
@apiParam {String} password 密码
@apiParam {String} nickname 昵称
@apiParam {String} phone_number 手机号码
@apiParam {String} internal_number 内部电话号码

@apiError UserExists The user account already exists
@apiErrorExample Error-Response:
HTTP/1.1 400 Bad Request
{
   "errorCode": 10011,
   "errorMessage": "user_exists"
}
*/
func (*Servlet) CreateUser(c *gin.Context) {

	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidArg.WithMessage(err.Error()))
		return
	}

	ctx := c.Request.Context()

	userWithName, err := model.UserStatic.GetByName(ctx, req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if userWithName != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgUserExists)
		return
	}

	userWithPhone, err := model.UserStatic.GetByPhone(ctx, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}
	if userWithPhone != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgUserExists)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if err := model.UserStatic.Insert(ctx, &model.User{
		UserName:            req.UserName,
		Password:            hashedPassword,
		PhoneNumber:         req.PhoneNumber,
		Nickname:            req.Nickname,
		InternalPhoneNumber: req.InternalPhone,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, statusmsg.StatusSuccess)
}

/**
@api {post} /venus/auth/user/signin_with_name 用户名登录
@apiGroup User

@apiParam {String} user_name 用户名
@apiParam {String} password 密码

@apiSuccess {String} token Token.

@apiError UserNotFound User not found
@apiErrorExample Error-Response:
HTTP/1.1 404 Not Found
{
    "errorCode": 10010,
    "errorMessage": "user_not_found"
}

@apiError InvalidOAuthCode Invalid Token
@apiErrorExample Error-Response:
HTTP/1.1 404 Bad Request
{
    "errorCode": 11003,
    "errorMessage": "invalid_oauth_code"
}
*/
func (*Servlet) SigninUserByName(c *gin.Context) {

	var req signinUserByNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidArg.WithMessage(err.Error()))
		return
	}

	ctx := c.Request.Context()

	user, err := model.UserStatic.GetByName(ctx, req.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidOAuthCode)
		return
	}

	tokenString, err := util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if err := model.UserStatic.UpdateLoginTime(ctx, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &signinUserResp{
		Token: tokenString,
	})
}

/**
@api {post} /venus/auth/user/signin_with_phone 手机号码登录
@apiGroup User

@apiParam {String} phone_number 手机号码
@apiParam {String} password 密码

@apiSuccess {String} token Token

@apiErrorExample Error-Response:
HTTP/1.1 403 Bad Request
{
    "errorCode": 11003,
    "errorMessage": "invalid_oauth_code"
}
*/
func (*Servlet) SigninUserByPhone(c *gin.Context) {

	var req signinUserByPhoneReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidArg.WithMessage(err.Error()))
		return
	}

	ctx := c.Request.Context()

	user, err := model.UserStatic.GetByPhone(ctx, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidOAuthCode)
		return
	}

	tokenString, err := util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if err = model.UserStatic.UpdateLoginTime(ctx, user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &signinUserResp{
		Token: tokenString,
	})
}

/**
@api {post} /venus/user/info 更新手机号码/昵称
@apiGroup User
@apiHeader {String} token Token
@apiParam {String} phone_number 新手机号码
@apiParam {String} nickname 用户昵称
@apiSuccess {String} success 状态码

@apiSuccessExample Success-Response:
HTTP/1.2 200 OK
{
 "success": "ok"
}
*/
func (*Servlet) UpdateUserPhoneNumberAndNickname(c *gin.Context) {

	var req updatePhoneNumberAndNicknameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidArg.WithMessage(err.Error()))
		return
	}

	userID, _ := auth.GetUserIdAndNickname(c)
	ctx := c.Request.Context()

	if err := model.UserStatic.UpdateUserInfo(ctx, userID, req.Nickname, req.PhoneNumber); err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, statusmsg.StatusSuccess)
}

/**
@api {post} /venus/token/refresh 获取refresh token
@apiGroup User

@apiHeader {String} token Token
HTTP/1.2 200 OK
{
 "token": "token"
}
*/
func (*Servlet) GetRefreshToken(c *gin.Context) {

	ctx := c.Request.Context()
	userID, _ := auth.GetUserIdAndNickname(c)

	user, err := model.UserStatic.GetByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	tokenString, err := util.GenerateToken(user.ID, user.Nickname, user.RoleID, global.Config.JWTTokenExpSec, global.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &signinUserResp{
		Token: tokenString,
	})
}

/**
@api {post} /venus/user/password 用户更新个人密码
@apiGroup User
@apiHeader {String} token Token
@apiParam {String} password 用户密码
@apiParam {String} confirm_password 用户确认密码

@apiSuccess {String} success 状态码

@apiSuccessExample Success-Response:
HTTP/1.2 200 OK
{
 "success": "ok"
}
*/
func (*Servlet) EditUserPassword(c *gin.Context) {
	var req updateUserPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgInvalidArg.WithMessage(err.Error()))
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, statusmsg.StatusMsgUserPasswordNotEqual)
		return
	}
	ctx := c.Request.Context()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	userID, _ := auth.GetUserIdAndNickname(c)
	if userID == 0 {
		c.JSON(http.StatusNotFound, statusmsg.StatusMsgUserNotFound)
		return
	}

	err = model.UserStatic.UpdateUserPassword(ctx, userID, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, statusmsg.StatusMsgProcessError.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, statusmsg.StatusSuccess)
}
