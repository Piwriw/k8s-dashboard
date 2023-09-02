package request

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type NfsVolumeSource struct {
	NfPath      string `json:"nfPath"`
	NfsServer   string `json:"nfsServer"`
	NfsReadOnly bool   `json:"nfsReadOnly"`
}
type VolumeSource struct {
	Type            string          `json:"type"`
	NfsVolumeSource NfsVolumeSource `json:"nfsVolumeSource"`
}
type PersistentVolume struct {
	Name string `json:"name"`
	//Namespace string             `json:"namespace"`
	Labels []base.ListMapItem `json:"labels"`
	//pv 容量
	Capacity int32 `json:"capacity"`
	//读写权限
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	//pv回收策略
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
	VolumeSource  VolumeSource                         `json:"volumeSource"`
}
