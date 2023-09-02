package pv

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/model/base"
	pv_req "k8sdashboar.com/model/pv/request"
	"strconv"
)

type PVReq2K8sConvert struct {
}

func (pv *PVReq2K8sConvert) PVReq2K8s(pvReq pv_req.PersistentVolume) *corev1.PersistentVolume {
	return &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvReq.Name,
			Labels: base.ToMap(pvReq.Labels),
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(pvReq.Capacity)) + "Mi"),
			},
			AccessModes:                   pvReq.AccessModes,
			PersistentVolumeReclaimPolicy: pvReq.ReClaimPolicy,
			PersistentVolumeSource:        pv.GetK8sPVVolumeSource(pvReq.VolumeSource),
		},
	}
}
func (*PVReq2K8sConvert) GetK8sPVVolumeSource(pvVolumeSourceReq pv_req.VolumeSource) corev1.PersistentVolumeSource {
	var volumeSource corev1.PersistentVolumeSource
	switch pvVolumeSourceReq.Type {
	case "nfs":
		volumeSource.NFS = &corev1.NFSVolumeSource{
			Server:   pvVolumeSourceReq.NfsVolumeSource.NfsServer,
			Path:     pvVolumeSourceReq.NfsVolumeSource.NfPath,
			ReadOnly: pvVolumeSourceReq.NfsVolumeSource.NfsReadOnly,
		}
	default:
		//response.FailWithMessage(c, "不支持的存储卷")
		return volumeSource
	}
	return volumeSource
}
