package request

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type UpdateLabel struct {
	Name   string             `json:"name"`
	Labels []base.ListMapItem `json:"labels"`
}
type UpdateTaint struct {
	Name   string         `json:"name"`
	Taints []corev1.Taint `json:"taints"`
}
