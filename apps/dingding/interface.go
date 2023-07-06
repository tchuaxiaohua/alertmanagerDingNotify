package dingding

import (
	"fmt"

	"github.com/tchuaxiaohua/alertmanagerDingNotify/config"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	Send(data string)
}

// Send 告警发送
func (d *DingDing) Send(data string) {
	dingClient := dingtalk.NewClient(d.Token, "")
	msg := dingtalk.NewMarkdownMessage().SetMarkdown(d.Title, data)
	dingClient.Send(msg)
}

func (d *DingDing) GetToken(c *gin.Context) {
	// 获取用户渠道名称和token值
	appName := c.Param("app")
	if appName == "" {
		zap.L().Error("请求地址错误", zap.String("message", "无法获取渠道名称"))
		return
	}
	appToken, ok := config.C().Ding[appName]
	if !ok {
		zap.L().Error("获取token失败", zap.String("message", "配置文件token获取失败，请检查对应渠道token配置"))
		return
	}
	zap.L().Debug("应用和token", zap.String("message", fmt.Sprintf("告警渠道:%s", appName)))

	d.Token = appToken
	d.AppName = appName
}
