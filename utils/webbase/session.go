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

type MonitorUserCtx struct {
	UserId   string `json:"UID"`
	Username string `json:"UNM"`
	RoleId   int    `json:"RID"`
}

const (
	LoginStatusKey = "login.status"
	UserLoginKey   = "login.user"

	MonitorLoginStatusKey = "login.monitor.status"
	MonitorLoginKey       = "login.monitor.user"
)

const (
	KUserLogin = 1
)

func GetMonitorUserCtx(c *gin.Context) *MonitorUserCtx {
	return c.MustGet(MonitorLoginKey).(*MonitorUserCtx)
}

func GetUserCtx(c *gin.Context) *UserCtx {
	return c.MustGet(UserLoginKey).(*UserCtx)
}

func (ctx *UserCtx) GetUserId() string {
	return ctx.UserId
}

func (ctx *UserCtx) IsLogin() bool {
	return ctx.UserId != ""
}
