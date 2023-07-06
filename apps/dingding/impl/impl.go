package impl

import (
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/dingding"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/k8s"
	"github.com/tchuaxiaohua/alertmanagerDingNotify/apps/prometheus"
)

type DingService struct {
	dSvc dingding.Service
	pSvc prometheus.AlertManager
	kSvc k8s.K8s
}
