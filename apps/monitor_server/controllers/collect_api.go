package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mongo"
	"shopping/utils/webbase"

	"shopping/apps/monitor_server/models"
)

// CollectData 收集监控数据
func CollectData(c *gin.Context) {
	req := &mongo.MonitorData{}
	if err := c.ShouldBindJSON(req); err != nil {
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	if err := models.InsertMonitorData(req); err != nil {
		Logger.Warn("InsertMonitorData err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	webbase.ServeResponse(c, webbase.ErrOK)
}
