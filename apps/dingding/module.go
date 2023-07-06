package dingding

import (
	"github.com/tchuaxiaohua/alertmanagerDingNodify/apps/prometheus"
)

type DingDing struct {
	Title       string // 告警标题
	AppName     string // 告警渠道
	Token       string // 告警渠道钉钉Token
	AlertManage *prometheus.AlertManager
}

func NewDingDing(title string) *DingDing {
	return &DingDing{
		Title:       title,
		AlertManage: prometheus.NewAlertManager(),
	}
}
