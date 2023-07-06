package http

import (
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/dingding"
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
	for _, alertMsg := range ins.AlertManage.Alerts {
		if err := alertMsg.TimeFormat(); err != nil {
			zap.L().Error("告警时间解析失败", zap.String("error", err.Error()), zap.String("message", "时间解析出错"))
			return
		}

		// 判断是否需要执行dump操作
		//_, ok := alertMsg.Labels["jvm_dump"]
		//if ok {
		//	if config.C().Jvm.IsDump {
		//		// TODO: 操作k8s dump
		//		err := alertMsg.Dump()
		//		if err != nil {
		//			zap.L().Warn("jvm dump", zap.String("message", "执行dump操作失败"), zap.String("error", err.Error()))
		//			continue
		//		}
		//		// TODO: 触发dump 操作发送钉钉通知
		//		zap.L().Info("jvm dump", zap.String("message", "dump成功，发送钉钉通知"))
		//		continue
		//	}
		//}
		//data, err := utils.ParseTemplate("template/alert.tmpl", alertMsg)
		//
		////data, err := alertMsg.ParseAlertTemplate("template/alert.tmpl")
		//if err != nil {
		//	zap.L().Error("模板解析失败", zap.String("error", err.Error()), zap.String("message", "解析钉钉告警模板失败"))
		//	return
		//}
		data := ins.AlertManage.AlarmRoute()
		ins.Send(data)
	}

}
