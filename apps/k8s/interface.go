package k8s

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (k *K8s) Exec(cmd []string) error {
	// su admin -c "jmap -dump:live,format=b,file=/tmp/dump.hprof pid"
	//cmd := []string{"sh", "-c", "jmap -dump:live,format=b,file=/tmp/${HOSTNAME}.hprof 1"}
	req := k.ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(k.PodName).Namespace(k.NameSpace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
			Command: cmd,
		}, scheme.ParameterCodec)

	// 创建执行器
	exec, err := remotecommand.NewSPDYExecutor(k.Config, "POST", req.URL())
	if err != nil {
		fmt.Println("创建执行器失败:", err)
		return err
	}
	var stdout, stderr bytes.Buffer
	err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Output from pod: %v", stdout.String())
	fmt.Printf("Error from pod: %v", stderr.String())
	return nil
}

func (k *K8s) GetPod() {
	pod, err := k.ClientSet.CoreV1().Pods(k.NameSpace).Get(context.TODO(), k.PodName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error("GetPod failed", zap.String("message", "根据pod名称获取podIP失败"), zap.String("error", err.Error()))
		return
	}
	k.PodIP = pod.Status.PodIP
}

func (k *K8s) ListEvents() {
	podEvents, _ := k.ClientSet.CoreV1().Events(k.NameSpace).List(context.TODO(), metav1.ListOptions{})
	for _, v := range podEvents.Items {
		if v.Type == "Normal" {
			continue
		}

		if strings.HasPrefix(v.Name, k.PodName) {
			k.Events = append(k.Events, v.Message)
		}
	}
}
