package main

import (
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/dingding/http"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/config"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化 全局配置文件
	if err := config.LoadConfig("etc/app.yaml"); err != nil {
		zap.L().Error("加载配置失败", zap.String("err", err.Error()))
	}
	// 初始化日志
	config.InitLogger("info")

	g := gin.New()
	g.Use(middleware.GinLogger(config.Log), middleware.GinRecovery(config.Log, true))
	g.GET("/health/", health())
	g.POST("/api/dingding/:app", http.SendDingTalk)
	g.Run("0.0.0.0:18084")
}

func health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "hello dingtalk",
		})
	}
}
