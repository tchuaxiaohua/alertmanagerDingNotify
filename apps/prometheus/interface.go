package prometheus

import (
	"errors"
	"fmt"
	"time"

	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/k8s"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/config"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/utils"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/vars"

	"go.uber.org/zap"
)

// TimeFormat 解析告警时间为标准时间
func (a *Alert) TimeFormat() error {
	layout := "2006-01-02 15:04:05"

	// 告警触发时间
	t, err := time.Parse(time.RFC3339, a.StartsAt)
	if err != nil {
		return err
	}
	a.StartsAt = t.In(time.Local).Format(layout)

	// 告警结束时间
	if len(a.EndsAt) > 0 {
		t1, err := time.Parse(time.RFC3339, a.EndsAt)
		if err != nil {
			return err
		}
		a.EndsAt = t1.In(time.Local).Format(layout)
	}
	return nil
}

// Dump jvm内存快照
func (a *Alert) Dump() error {
	k8sClient, err := k8s.NewK8s()
	if err != nil {
		zap.L().Error("Failed to create k8s client", zap.String("message", "调用下游函数失败"))
		return err
	}
	// 处理所需要的标签值
	currentValue := a.Annotations.CurrentValue
	k8sClient.PodName = a.getMap("pod")
	k8sClient.NameSpace = a.getMap("namespace")

	if a.Status == "resolved" {
		return errors.New(fmt.Sprintf("【%s】告警恢复，无需处理", k8sClient.PodName))
	}

	// 判断是否超过一天时间 超过一天才需要dump 这里需要设置对应的repeat_interval
	if time.Since(utils.PtrTime(a.StartsAt)) > time.Duration(time.Duration(config.C().Jvm.DumpTsMax)*time.Hour) {
		return errors.New(fmt.Sprintf("【%s】告警时间超过%dH,已执行dump操作,告警触发时间:%s", k8sClient.PodName, config.C().Jvm.DumpTsMax, a.StartsAt))
	}

	if !(time.Since(utils.PtrTime(a.StartsAt)) > time.Duration(time.Duration(config.C().Jvm.DumpTsMin)*time.Hour)) {
		return errors.New(fmt.Sprintf("【%s】告警触发时间不足%dH,告警触发时间:%s", k8sClient.PodName, config.C().Jvm.DumpTsMin, a.StartsAt))
	}

	if utils.PtrInt(currentValue) >= config.C().Jvm.DumpMin && utils.PtrInt(currentValue) <= config.C().Jvm.DumpMax {
		// 执行jvm内存快照导出
		if err := k8sClient.Exec(vars.CmdDump); err != nil {
			return err
		}
		// TODO: 告警通知 直接调用dump文件上传脚本
		// 上传快展文件至oss 并通知
		if err := k8sClient.Exec(vars.CmdUploadDump); err != nil {
			return err
		}
		// ---------- 集成oss upload ------------------------
		//_, err := utils.ParseTemplate("etc/jvmDump.tmpl", "")
		//if err != nil {
		//	zap.L().Error("模板dump解析失败", zap.String("error", err.Error()), zap.String("message", "解析钉钉告警模板失败"))
		//	return err
		//}
		//a.ParseAlertTemplate("etc/jvmDump.tmpl")
		//data, err := a.ParseAlertTemplate("etc/jvmDump.tmpl")
		//if err != nil {
		//	zap.L().Error("模板解析失败", zap.String("error", err.Error()), zap.String("message", "解析钉钉告警模板失败"))
		//	return err
		//}
		return nil
	} else {
		return errors.New(fmt.Sprintf("【%s】当前值内存使用率已超过%d%%,当前值:%d%%,退出执行dump操作！！！", k8sClient.PodName, config.C().Jvm.DumpMax, utils.PtrInt(currentValue)))
	}
}

// getMap 从Labels 中取值
func (a *Alert) getMap(key string) string {
	v, ok := a.Labels[key]
	if ok {
		return v.(string)
	}
	return ""
}

func (a *Alert) getEvents() {
	k8sClient, err := k8s.NewK8s()
	if err != nil {
		zap.L().Error("Failed to create k8s client", zap.String("message", "调用下游函数失败"))
	}
	// 取podEvents
	k8sClient.PodName = a.getMap("pod")
	k8sClient.NameSpace = a.getMap("namespace")
	k8sClient.ListEvents()
	a.Events = k8sClient.Events
}

func (a *AlertManager) AlarmRoute() string {
	// 数据解析 如果有多条告警信息 则分批处理发送
	for _, alertMsg := range a.Alerts {
		if err := alertMsg.TimeFormat(); err != nil {
			zap.L().Error("告警时间解析失败", zap.String("error", err.Error()), zap.String("message", "时间解析出错"))
			return ""
		}

		// 判断是否需要执行dump操作
		_, ok := alertMsg.Labels["jvm_dump"]
		if ok {
			if config.C().Jvm.IsDump {
				// TODO: 操作k8s dump
				err := alertMsg.Dump()
				if err != nil {
					zap.L().Warn("jvm dump", zap.String("message", "执行dump操作失败"), zap.String("error", err.Error()))
					continue
				}
				// TODO: 触发dump 操作发送钉钉通知
				zap.L().Info("jvm dump", zap.String("message", "dump成功，发送钉钉通知"))
				continue
			}
		}

		// pod 事件处理 只针对pod告警处理
		_, ok = alertMsg.Labels["pod"]
		if ok {
			alertMsg.getEvents()
		}
		// 数据解析
		data, err := utils.ParseTemplate("template/alert.tmpl", alertMsg)
		//data, err := alertMsg.ParseAlertTemplate("template/alert.tmpl")
		if err != nil {
			zap.L().Error("模板解析失败", zap.String("error", err.Error()), zap.String("message", "解析钉钉告警模板失败"))
			return ""
		}
		return data
	}
	return ""
}
