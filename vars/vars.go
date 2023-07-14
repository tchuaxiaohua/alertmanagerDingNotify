package vars

const (
	PrometheusK8SWarningTitle = "Prometheus告警通知"
	PrometheusJvmDumpTitle    = "应用JVM Dump通知"
)

var (
	// 注释: 如果应用在pod中进程 1 则可以直接使用这个命令 执行导出即可
	//CmdDump       = []string{"sh", "-c", "jmap -dump:live,format=b,file=/tmp/${HOSTNAME}.hprof 1"}
	// 注释: 这里 配置一个dump.sh脚本，是为了适配应用进程为非 1 的时候
	// 脚本内容: https://wiki.tbchip.com/pages/821ce3/#_2-4-dump-%E9%85%8D%E7%BD%AE
	CmdDump       = []string{"sh", "-c", "/devops/dump.sh"}
	CmdUploadDump = []string{"sh", "-c", "/devops/cloud-station  -f /tmp/${HOSTNAME}.hprof"}
	//CmdShell      = []string{"sh", "-c", "/devops/dump.sh"}
)
