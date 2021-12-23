package controllers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/monitor_server/io"
	"shopping/apps/monitor_server/models"
)

// Login 登录
func Login(c *gin.Context) {
	req := &io.LoginRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(200, webbase.ErrInputParams)
		return
	}

	session := sessions.Default(c)
	iface := session.Get(webbase.UserLoginKey)

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

	// 将request中的密码加密后，再比较
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		Logger.Warn("bcrypt generate from password err", zap.Error(err))
		c.JSON(200, webbase.ErrSystemBusy)
		return
	}

	if user.Password != string(hash) {
		Logger.Warn("password err")
		c.JSON(200, ErrPassword)
		return
	}

	su := &webbase.MonitorUserCtx{
		UserId:   user.Id,
		Username: user.Username,
		RoleId:   user.RoleId,
	}
	c.Set(webbase.MonitorLoginKey, su)
	session.Set(webbase.MonitorLoginStatusKey, webbase.MonitorLoginStatusKey)
	session.Set(webbase.MonitorLoginKey, su)
	session.Save()

	c.JSON(200, webbase.ErrOK)
}
