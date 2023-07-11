package vars

const (
	PrometheusK8SWarningTitle = "Prometheus告警通知"
	PrometheusJvmDumpTitle    = "应用JVM Dump通知"
)

var (
	//CmdDump       = []string{"sh", "-c", "jmap -dump:live,format=b,file=/tmp/${HOSTNAME}.hprof 1"}
	CmdDump       = []string{"sh", "-c", "/devops/dump.sh"}
	CmdGetPodIP   = []string{"sh", "-c", "echo ${KUBERNETES_POD_IP}"}
	CmdUploadDump = []string{"sh", "-c", "/devops/cloud-station  -f /tmp/${HOSTNAME}.hprof"}
	//CmdShell      = []string{"sh", "-c", "/devops/dump.sh"}
)
