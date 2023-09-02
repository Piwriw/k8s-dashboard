package pvc

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/global"
	pvc_req "k8sdashboar.com/model/pvc/request"
	pvc_res "k8sdashboar.com/model/pvc/response"
	"strings"
)

type PVCService struct {
}

// DeletePVC 根据name删除PVC
func (*PVCService) DeletePVC(name, namespace string) error {
	return global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

// GetPVCList 获取PVC列表
func (*PVCService) GetPVCList(keyword, namespace string) ([]pvc_res.PVCItem, error) {
	k8sPVCList, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	pvcResList := make([]pvc_res.PVCItem, 0)
	for _, item := range k8sPVCList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		pvcResList = append(pvcResList, pvcConvert.GetPVItemRes(item))
	}

	return pvcResList, nil
}

// GetPVCDetail 查询PVC详情
func (*PVCService) GetPVCDetail(name, namespace string) (*corev1.PersistentVolumeClaim, error) {
	k8sPVC, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return k8sPVC, nil
}

// CreatePVC 创建PVC
func (*PVCService) CreatePVC(pvcReq pvc_req.PersistentVolumeClaim) error {

	pvcK8s := pvcConvert.PVCReq2K8s(pvcReq)
	_, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(pvcReq.Namespace).Create(context.TODO(), pvcK8s, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
