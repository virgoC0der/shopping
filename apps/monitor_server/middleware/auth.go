package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"shopping/utils/mysql/monitor"
	"strings"

	"github.com/gin-gonic/gin"

	. "shopping/utils/log"
	"shopping/utils/webbase"
)

var appCodeMap = map[string]struct{}{
	"monitor_agent": {},
}

const (
	kAuthSliceLen = 2
)

func CheckAuth(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	authSlice := strings.Split(auth, " ")
	if len(authSlice) < kAuthSliceLen {
		Logger.Warn("auth header is invalid", zap.String("Authorization", auth))
		c.AbortWithStatusJSON(http.StatusOK, webbase.ErrAuthFailed)
		return
	}

	if _, ok := appCodeMap[authSlice[kAuthSliceLen-1]]; !ok {
		Logger.Warn("app code is invalid", zap.String("appCode", authSlice[kAuthSliceLen-1]))
		c.AbortWithStatusJSON(http.StatusOK, webbase.ErrAuthFailed)
		return
	}

	c.Next()
}

func CheckPermission(c *gin.Context) {
	ctx := webbase.GetMonitorUserCtx(c)
	if ctx == nil {
		Logger.Warn("user has not login")
		c.AbortWithStatusJSON(http.StatusOK, webbase.ErrNotLogin)
		return
	}

	if ctx.RoleId != monitor.KAdminRoleId {
		Logger.Warn("user has no permission", zap.String("user_id", ctx.UserId))
		c.AbortWithStatusJSON(http.StatusOK, webbase.ErrNoPermission)
		return
	}

	c.Next()
}
