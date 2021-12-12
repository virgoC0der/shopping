package webbase

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserCtx struct {
	UserId   string `json:"UID"`
	Username string `json:"UNM"`
	Phone    string `json:"PHO"`
}

const (
	LoginStatusKey = "login.status"
	UserLoginKey   = "login.user"
)

const (
	KUserLogin = 1
)

func (ctx *UserCtx) GetUserCtx(c *gin.Context) *UserCtx {
	session := sessions.Default(c)
	iface := session.Get(UserLoginKey)
	userSession, ok := iface.(string)
	if !ok {
		return nil
	}

	err := json.Unmarshal([]byte(userSession), ctx)
	if err != nil {
		return nil
	}

	return ctx
}

func (ctx *UserCtx) SetUserCtx(userCtx *UserCtx) {
	ctx.UserId = userCtx.UserId
	ctx.Username = userCtx.Username
	ctx.Phone = userCtx.Phone
}

func (ctx *UserCtx) GetUserId() string {
	return ctx.UserId
}

func (ctx *UserCtx) IsLogin() bool {
	return ctx.UserId != ""
}
