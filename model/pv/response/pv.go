package response

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type PVItem struct {
	Name string `json:"name"`
	//Namespace string             `json:"namespace"`
	Labels []base.ListMapItem `json:"labels"`
	//pv 容量
	Capacity int32 `json:"capacity"`
	//读写权限
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	//pv回收策略
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
	Status        corev1.PersistentVolumePhase         `json:"status"`
	//哪个pvc绑定
	Claim string `json:"claim"`
	Age   int64  `json:"age"`
	//状态描述
	Reason           string `json:"reason"`
	StorageClassName string `json:"storageClassName"`
}
