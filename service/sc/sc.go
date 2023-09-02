package sc

import (
	"context"
	"errors"
	"fmt"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/global"
	sc_req "k8sdashboar.com/model/sc/request"
	sc_res "k8sdashboar.com/model/sc/response"
	"strings"
)

const (
	NFS_SUBDIR = "cluster.local/nfs-subdir-external-provisioner"
)

type SCService struct {
}

// DeleteSC 根据name删除SC
func (*SCService) DeleteSC(name string) error {
	return global.KubeConfigSet.StorageV1().StorageClasses().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

// GetSCList 获取SC列表
func (*SCService) GetSCList(keyword string) ([]sc_res.StorageClassItem, error) {
	k8sSCList, err := global.KubeConfigSet.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	scResList := make([]sc_res.StorageClassItem, 0)
	for _, item := range k8sSCList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		scResList = append(scResList, scConvert.GetSCItemRes(item))
	}

	return scResList, nil
}

// GetSCDetail 查询SC详情
func (*SCService) GetSCDetail(name string) (*storagev1.StorageClass, error) {
	k8sSC, err := global.KubeConfigSet.StorageV1().StorageClasses().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return k8sSC, nil
}

// CreateSC 创建SC
func (*SCService) CreateSC(scReq sc_req.StorageClass) error {
	provisionerList := strings.Split(global.CONF.System.Provisioner, ",")
	var flag bool
	for _, val := range provisionerList {
		if scReq.Provisioner == val {
			flag = true
		}
	}
	if !flag {
		errMsg := fmt.Sprintf("%s 当前K8S未支持", scReq.Provisioner)
		return errors.New(errMsg)
	}

	scK8s := scConvert.SCReq2K8s(scReq)
	_, err := global.KubeConfigSet.StorageV1().StorageClasses().Create(context.TODO(), scK8s, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
