package prometheus

// AlertManager 告警信息
type AlertManager struct {
	Receiver string  `json:"receiver"`
	Status   string  `json:"status"`
	Alerts   []Alert `json:"alerts"`
}

func NewAlertManager() *AlertManager {
	return &AlertManager{}
}

// Alert 具体告警内容
type Alert struct {
	Status       string                 `json:"status"`
	Labels       map[string]interface{} `json:"labels"`
	Annotations  Annotations            `json:"annotations"`
	StartsAt     string                 `json:"startsAt"`
	EndsAt       string                 `json:"endsAt"`
	GeneratorUrl string                 `json:"generatorURL"`
	FingerPrint  string                 `json:"fingerprint"`
	Events       []string
}

// Annotations 描述
type Annotations struct {
	Description  string `json:"description"`
	Summary      string `json:"summary"`
	CurrentValue string `json:"currentvalue"`
}

// DumpData 解析使用
type DumpData struct {
	PodName   string `json:"podName"`
	PodIP     string `json:"podIP"`
	Title     string
	CreatedAt string `json:"createdAt"`
	Url       string `json:"url"`
}
