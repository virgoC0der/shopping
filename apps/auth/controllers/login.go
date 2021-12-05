package controllers

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/auth/io"
	"shopping/apps/auth/models"
)

// Login 登录
func Login(c *gin.Context) {
	req := &io.LoginRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(200, webbase.ErrInputParams)
		return
	}

	session := sessions.Default(c)
	iface := session.Get("user")

	userSession, ok := iface.(string)
	if !ok {
		userSession = ""
	}

	ctx := &webbase.UserCtx{}
	if err := json.Unmarshal([]byte(userSession), ctx); err != nil {
		Logger.Warn("json unmarshal user session err", zap.Error(err))
	}

	user, err := models.GetUser(req.Username)
	if err != nil {
		Logger.Warn("get user err", zap.Error(err))
		c.JSON(200, webbase.ErrSystemBusy)
		return
	}

	if user.Password != req.Password {
		Logger.Warn("password err")
		c.JSON(200, ErrPassword)
		return
	}

	su := &webbase.UserCtx{
		UserId:   user.Id,
		Username: user.Username,
		Phone:    user.Phone,
	}
	session.Set(webbase.LoginStatusKey, webbase.KUserLogin)
	session.Set(webbase.UserLoginKey, su)
	session.Save()

	c.JSON(200, webbase.ErrOK)
}
