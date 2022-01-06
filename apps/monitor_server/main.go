package main

import (
	"encoding/gob"
	"go.uber.org/zap"
	"os"
	"shopping/utils/mongo"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	. "shopping/utils/log"
	"shopping/utils/valid"
	"shopping/utils/webbase"

	"shopping/apps/monitor_server/controllers"
	"shopping/apps/monitor_server/middleware"
	"shopping/apps/monitor_server/middleware/prometheus"
)

func main() {
	InitLog()

	if err := mongo.Init(); err != nil {
		Logger.Error("init mongo err", zap.Error(err))
		os.Exit(1)
	}

	r := gin.New()
	// 初始化session
	gob.Register(webbase.MonitorUserCtx{})
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions(webbase.UserLoginKey, store))
	r.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(Logger, true))
	gp := prometheus.New(r)
	r.Use(gp.Middleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid.Register(v)
	}

	// 收集数据api
	agent := r.Group("/montior/api/v1")
	{
		agent.Use(middleware.CheckAuth)
		agent.POST("/collect", controllers.CollectData)
	}

	// web api
	monitor := r.Group("/monitor/v1")
	{
		monitor.POST("/login", controllers.Login)
		monitor.POST("/data", controllers.GetMonitorData)
	}

	admin := r.Group("/monitor/v1/admin")
	{
		admin.Use(middleware.CheckPermission)
		admin.POST("/log", controllers.GetUserLog)
	}

	// prometheus metrics collect api
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8081")
}
