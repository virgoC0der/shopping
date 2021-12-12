package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"shopping/apps/auth/models"
	. "shopping/utils/log"
	"shopping/utils/webbase"
)

func GetUserInfo(c *gin.Context) {
	ctx := webbase.GetUserCtx(c)
	user, err := models.GetUserById(ctx.UserId)
	if err != nil {
		Logger.Warn("get user by id err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	webbase.ServeResponse(c, webbase.ErrOK, user)
}
