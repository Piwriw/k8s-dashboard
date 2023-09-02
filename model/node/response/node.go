package response

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type Node struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Age        int64  `json:"age"`
	InternalIP string `json:"internalIp"`
	ExternalIP string `json:"externalIP"`
	// kubectl version
	Version       string `json:"version"`
	OsImage       string `json:"osImage"`
	KernelVersion string `json:"kernelVersion"`
	// 容器运行时
	ContainerRuntime string             `json:"containerRuntime"`
	Labels           []base.ListMapItem `json:"labels"`
	Taints           []corev1.Taint     `json:"taints"`
}
