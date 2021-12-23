package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"shopping/utils/mongo"
	"time"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/monitor_server/io"
	"shopping/apps/monitor_server/models"
)

// GetMonitorData 查询监控数据
func GetMonitorData(c *gin.Context) {
	req := &io.GetMonitorDataReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	// 记录操作日志
	defer func() {
		ctx := webbase.GetMonitorUserCtx(c)
		log := &mongo.UserLog{
			UserId:      ctx.UserId,
			Time:        time.Now().Unix(),
			QueryParams: req,
		}
		if err := models.InsertUserLog(log); err != nil {
			Logger.Warn("insert user log error", zap.Error(err))
		}
	}()

	filter := bson.M{
		"time": bson.M{
			"$gte": req.Start,
			"$lte": req.End,
		},
	}
	if len(req.System) > 0 {
		filter["system"] = bson.M{
			"$in": req.System,
		}
	}

	if req.AgentName != "" {
		filter["agent_name"] = req.AgentName
	}

	// 查询数据
	monitorData, err := models.QueryMonitorData(filter)
	if err != nil {
		Logger.Warn("query monitor data err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	webbase.ServeResponse(c, webbase.ErrOK, monitorData)
}
