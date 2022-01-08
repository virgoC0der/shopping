package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/monitor_server/io"
	"shopping/apps/monitor_server/models"
)

// GetUserLog 获取用户日志
func GetUserLog(c *gin.Context) {
	req := io.GetUserLogReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	offset := req.PageSize * req.PageSize
	total, logs, err := models.QueryUserLog(int64(req.PageSize), int64(offset))
	if err != nil {
		Logger.Warn("query user log err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	resp := io.GetUserLogResp{
		Total: total,
		Logs:  logs,
	}
	webbase.ServeResponse(c, webbase.ErrOK, resp)
}
