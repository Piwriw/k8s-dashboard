package pv

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/global"
	pv_req "k8sdashboar.com/model/pv/request"
	pv_res "k8sdashboar.com/model/pv/response"
	"strings"
)

type PVService struct {
}

func (*PVService) DeletePV(name string) error {
	return global.KubeConfigSet.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metav1.DeleteOptions{})
}
func (*PVService) GetPVList(keyword string) ([]pv_res.PVItem, error) {
	k8sPVList, err := global.KubeConfigSet.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	pvsResList := make([]pv_res.PVItem, 0)
	for _, item := range k8sPVList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		pvsResList = append(pvsResList, pvConvert.GetPVItemRes(item))
	}

	return pvsResList, nil
}
func (*PVService) GetPVDetail(name string) (*corev1.PersistentVolume, error) {
	k8sPV, err := global.KubeConfigSet.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return k8sPV, nil
}

func (*PVService) CreatePV(pvReq pv_req.PersistentVolume) error {

	pvK8s := pvConvert.PVReq2K8s(pvReq)
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumes().Create(context.TODO(), pvK8s, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
