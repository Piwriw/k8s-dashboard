package pvc

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	pvc_res "k8sdashboar.com/model/pvc/response"
)

type PVCK8s2Res struct {
}

func (pv *PVCK8s2Res) GetPVItemRes(k8sPVC corev1.PersistentVolumeClaim) pvc_res.PVCItem {
	matchLabels := make([]base.ListMapItem, 0)
	if k8sPVC.Spec.Selector != nil {
		matchLabels = base.ToList(k8sPVC.Spec.Selector.MatchLabels)
	}
	return pvc_res.PVCItem{
		Name:             k8sPVC.Name,
		Namespace:        k8sPVC.Namespace,
		Status:           k8sPVC.Status.Phase,
		Capacity:         int32(k8sPVC.Spec.Resources.Requests.Storage().Value() / (1024 * 1024)),
		AccessModes:      k8sPVC.Spec.AccessModes,
		StorageClassName: *k8sPVC.Spec.StorageClassName,
		Age:              k8sPVC.CreationTimestamp.Unix(),
		Volume:           k8sPVC.Spec.VolumeName,
		Labels:           base.ToList(k8sPVC.Labels),
		Selector:         matchLabels,
	}
}
func (*PVCK8s2Res) getPVItemClaim(k8sPV *corev1.ObjectReference) string {
	claim := ""
	if k8sPV != nil {
		claim = k8sPV.Name
	}
	return claim
}
