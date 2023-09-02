package request

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type PersistentVolumeClaim struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	//pv 容量
	Capacity         int32                               `json:"capacity"`
	AccessModes      []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	Selector         []base.ListMapItem                  `json:"selector"`
	StorageClassName string                              `json:"storageClassName"`
}
