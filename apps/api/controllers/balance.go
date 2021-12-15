package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"shopping/apps/api/io"
	"shopping/apps/api/models"
	. "shopping/utils/log"
	"shopping/utils/webbase"
)

func TopUpBalance(c *gin.Context) {
	req := &io.TopUpReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	userId := req.UserId
	if userId == "" {
		ctx := webbase.GetUserCtx(c)
		userId = ctx.UserId
	}

	if err := models.UpdateUserBalance(userId, req.Money); err != nil {
		Logger.Warn("update balance err", zap.Error(err), zap.String("userId", userId))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	webbase.ServeResponse(c, webbase.ErrOK)
}
