package k8s

import (
	"runtime"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8s struct {
	PodName   string
	NameSpace string
	PodIP     string
	Events    []string

	ClientSet kubernetes.Interface
	Config    *rest.Config
}

func NewK8s() (*K8s, error) {
	var err error
	var k8sConfig *rest.Config

	if runtime.GOOS == "windows" {
		k8sConfig, err = clientcmd.BuildConfigFromFlags("", "etc/config")
	} else {
		k8sConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		zap.L().Error("config failed", zap.String("error", err.Error()))
		return nil, err
	}
	// 初始化客户端
	k8sClient, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		zap.L().Error("client init failed", zap.String("error", err.Error()))
		return nil, err
	}
	return &K8s{
		ClientSet: k8sClient,
		Config:    k8sConfig,
	}, nil
}
