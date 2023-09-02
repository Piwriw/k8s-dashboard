package pv

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	pv_res "k8sdashboar.com/model/pv/response"
)

type PVK8s2Res struct {
}

func (pv *PVK8s2Res) GetPVItemRes(k8sPV corev1.PersistentVolume) pv_res.PVItem {
	return pv_res.PVItem{
		Name:             k8sPV.Name,
		Labels:           base.ToList(k8sPV.Labels),
		Capacity:         int32(k8sPV.Spec.Capacity.Storage().Value()) / (1024 * 1024),
		AccessModes:      k8sPV.Spec.AccessModes,
		ReClaimPolicy:    k8sPV.Spec.PersistentVolumeReclaimPolicy,
		Status:           k8sPV.Status.Phase,
		Claim:            pv.getPVItemClaim(k8sPV.Spec.ClaimRef),
		Age:              k8sPV.CreationTimestamp.Unix(),
		Reason:           k8sPV.Status.Reason,
		StorageClassName: k8sPV.Spec.StorageClassName,
	}
}
func (*PVK8s2Res) getPVItemClaim(k8sPV *corev1.ObjectReference) string {
	claim := ""
	if k8sPV != nil {
		claim = k8sPV.Name
	}
	return claim
}
