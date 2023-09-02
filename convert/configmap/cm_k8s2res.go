package configmap

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	configmap_res "k8sdashboar.com/model/configmap/response"
)

type CMK8s2Res struct {
}

func (*CMK8s2Res) GetCMDetailItem(cmK8s corev1.ConfigMap) configmap_res.ConfigMap {
	return configmap_res.ConfigMap{
		Name:      cmK8s.Name,
		Namespace: cmK8s.Namespace,
		DataNum:   len(cmK8s.Data),
		Age:       cmK8s.CreationTimestamp.Unix(),
	}
}
func (cm *CMK8s2Res) GetCMDetailRes(cmK8s corev1.ConfigMap) *configmap_res.ConfigMap {
	detail := cm.GetCMDetailItem(cmK8s)
	detail.Labels = base.ToList(cmK8s.Labels)
	detail.Data = base.ToList(cmK8s.Data)
	return &detail
}
