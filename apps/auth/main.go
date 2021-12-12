package main

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"shopping/apps/auth/controllers"
	. "shopping/utils/log"
	"shopping/utils/mysql"
	"shopping/utils/valid"
	"shopping/utils/webbase"
)

func main() {
	InitLog()

	if err := mysql.Init(); err != nil {
		Logger.Error("init mysql err", zap.Error(err))
		os.Exit(1)
	}

	r := gin.New()
	// 初始化session
	gob.Register(webbase.UserCtx{})
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions(webbase.UserLoginKey, store))
	r.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(Logger, true))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid.Register(v)
	}

	r.POST("/login", controllers.Login)
	r.GET("/product_list", controllers.GetProductList)
	r.GET("/product", controllers.GetProduct)
	r.GET("/user", controllers.GetUserInfo)
	r.Run(":8080")
}
