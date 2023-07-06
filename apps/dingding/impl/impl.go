package impl

import (
	"github.com/tchuaxiaohua/alertmanagerDingNodify/apps/dingding"
	"github.com/tchuaxiaohua/alertmanagerDingNodify/apps/k8s"
	"github.com/tchuaxiaohua/alertmanagerDingNodify/apps/prometheus"
)

type DingService struct {
	dSvc dingding.Service
	pSvc prometheus.AlertManager
	kSvc k8s.K8s
}
