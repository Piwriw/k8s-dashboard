package response

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type PVCItem struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	//读写权限
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	Capacity    int32                               `json:"capacity"`
	Selector    []base.ListMapItem                  `json:"selector"`
	Age         int64                               `json:"age"`
	//状态描述
	Reason           string                            `json:"reason"`
	StorageClassName string                            `json:"storageClassName"`
	Volume           string                            `json:"volume"`
	Status           corev1.PersistentVolumeClaimPhase `json:"status"`
}
