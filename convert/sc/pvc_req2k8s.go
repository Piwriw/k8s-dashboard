package sc

import (
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/model/base"
	sc_req "k8sdashboar.com/model/sc/request"
)

type SCReq2K8sConvert struct {
}

func (sc *SCReq2K8sConvert) SCReq2K8s(scReq sc_req.StorageClass) *storagev1.StorageClass {
	return &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:   scReq.Name,
			Labels: base.ToMap(scReq.Labels),
		},
		Provisioner:          scReq.Provisioner,
		MountOptions:         scReq.MountOptions,
		VolumeBindingMode:    &scReq.VolumeBindMode,
		ReclaimPolicy:        &scReq.ReclaimPolicy,
		AllowVolumeExpansion: &scReq.AllowVolumeExpansion,
		Parameters:           base.ToMap(scReq.Parameters),
	}
}
