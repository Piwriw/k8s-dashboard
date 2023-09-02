package sc

import (
	storagev1 "k8s.io/api/storage/v1"
	"k8sdashboar.com/model/base"
	sc_res "k8sdashboar.com/model/sc/response"
)

type SCK8s2Res struct {
}

func (pv *SCK8s2Res) GetSCItemRes(k8sSC storagev1.StorageClass) sc_res.StorageClassItem {
	var allowVolumeExpansion bool
	if k8sSC.AllowVolumeExpansion != nil {
		allowVolumeExpansion = *k8sSC.AllowVolumeExpansion
	}
	mountOptions := make([]string, 0)
	if k8sSC.MountOptions != nil {
		mountOptions = k8sSC.MountOptions
	}
	return sc_res.StorageClassItem{
		Name:                 k8sSC.Name,
		Labels:               base.ToList(k8sSC.Labels),
		Provisioner:          k8sSC.Provisioner,
		MountOptions:         mountOptions,
		Parameters:           base.ToList(k8sSC.Parameters),
		ReclaimPolicy:        *k8sSC.ReclaimPolicy,
		AllowVolumeExpansion: allowVolumeExpansion,
		VolumeBindMode:       *k8sSC.VolumeBindingMode,
		Age:                  k8sSC.CreationTimestamp.Unix(),
	}
}
