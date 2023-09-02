package pvc

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/model/base"
	pvc_req "k8sdashboar.com/model/pvc/request"
	"strconv"
)

type PVCReq2K8sConvert struct {
}

func (pv *PVCReq2K8sConvert) PVCReq2K8s(pvcReq pvc_req.PersistentVolumeClaim) *corev1.PersistentVolumeClaim {
	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcReq.Name,
			Namespace: pvcReq.Namespace,
			Labels:    base.ToMap(pvcReq.Labels),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: base.ToMap(pvcReq.Selector),
			},
			AccessModes: pvcReq.AccessModes,
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(pvcReq.Capacity)) + "Mi"),
				},
			},
			StorageClassName: &pvcReq.StorageClassName,
		},
	}
}
