package request

import (
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8sdashboar.com/model/base"
)

type StorageClass struct {
	Name string `json:"name"`
	//Namespace string             `json:"namespace"`
	Labels []base.ListMapItem `json:"labels"`
	// 制备器
	Provisioner string `json:"provisioner"`
	//  卷绑定参数配置
	MountOptions []string `json:"mountOptions"`
	// 制备器入参
	Parameters []base.ListMapItem `json:"parameters"`
	// 卷回收策略
	ReclaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	// 是否可拓展
	AllowVolumeExpansion bool `json:"allowVolumeExpansion"`
	// 卷绑定模式
	VolumeBindMode storagev1.VolumeBindingMode `json:"volumeBindMode"`
}
