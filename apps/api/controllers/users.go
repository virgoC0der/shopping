package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shopping/apps/api/io"
	"shopping/apps/api/models"
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

	orders, err := models.QueryOrders(ctx.UserId)
	if err != nil {
		Logger.Warn("query orders err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	resp := &io.GetUserInfoResp{
		Id:       user.Id,
		Username: user.Username,
		RealName: user.RealName,
		Phone:    user.Phone,
		Balance:  user.Balance,
		Orders:   orders,
	}
	webbase.ServeResponse(c, webbase.ErrOK, resp)
}
