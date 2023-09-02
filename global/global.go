package global

import (
	"k8s.io/client-go/kubernetes"
	"k8sdashboar.com/config"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
