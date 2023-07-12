package http

import (
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/dingding"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/prometheus"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/vars"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendDingTalk(c *gin.Context) {
	ins := dingding.NewDingDing(vars.PrometheusK8SWarningTitle)
	if err := c.ShouldBind(&ins.AlertManage); err != nil {
		zap.L().Error("Should Bind error", zap.String("error", err.Error()), zap.String("message", "参数解释错误"))
		return
	}

	// 获取用户渠道名称和token值
	ins.GetToken(c)

	// 数据解析 如果有多条告警信息 则分批处理发送
	zap.L().Info("")
	for _, alertMsg := range ins.AlertManage.Alerts {
		if err := alertMsg.TimeFormat(); err != nil {
			zap.L().Error("告警时间解析失败", zap.String("error", err.Error()), zap.String("message", "时间解析出错"))
			return
		}

		data := prometheus.AlertRoute(&alertMsg)
		// 返回空字符串 说明 解释失败 或者无需发送钉钉告警
		if data == "" {
			continue
		}
		ins.Send(data)
	}
}
