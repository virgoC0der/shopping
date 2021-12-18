package webbase

import (
	"github.com/gin-gonic/gin"
)

type UserCtx struct {
	UserId   string  `json:"UID"`
	Username string  `json:"UNM"`
	Nickname string  `json:"NKM"`
	Phone    string  `json:"PHO"`
	Balance  float64 `json:"BAL"`
}

const (
	LoginStatusKey = "login.status"
	UserLoginKey   = "login.user"
)

const (
	KUserLogin = 1
)

func GetUserCtx(c *gin.Context) *UserCtx {
	return c.MustGet(UserLoginKey).(*UserCtx)
}

func (ctx *UserCtx) GetUserId() string {
	return ctx.UserId
}

func (ctx *UserCtx) IsLogin() bool {
	return ctx.UserId != ""
}
