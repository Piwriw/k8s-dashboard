package configmap

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/global"
	configmap_req "k8sdashboar.com/model/configmap/request"
	configmap_res "k8sdashboar.com/model/configmap/response"
	"strings"
)

type ConfigMapService struct {
}

func (*ConfigMapService) GetConfigMapDetail(namespace, name string) (*configmap_res.ConfigMap, error) {
	k8sCm, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cmConvert.GetCMDetailRes(*k8sCm), nil
}
func (*ConfigMapService) GetConfigMapList(namespace, keyword string) ([]configmap_res.ConfigMap, error) {
	configMapList, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	cmListRes := make([]configmap_res.ConfigMap, 0)
	for _, item := range configMapList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		cmListRes = append(cmListRes, cmConvert.GetCMDetailItem(item))
	}
	return cmListRes, nil
}

// DeleteConfigMap 删除ConfigMap
func (*ConfigMapService) DeleteConfigMap(namespace, name string) error {
	ctx := context.TODO()
	return global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// CreateOrUpdateConfigMap 创建或者更新ConfigMap
func (*ConfigMapService) CreateOrUpdateConfigMap(configMapReq configmap_req.ConfigMap) error {
	ctx := context.TODO()
	cmK8s := cmConvert.CmReq2K8s(configMapReq)

	_, err := global.KubeConfigSet.CoreV1().ConfigMaps(configMapReq.Name).Get(ctx, configMapReq.Name, metav1.GetOptions{})
	if err == nil {
		_, err = global.KubeConfigSet.CoreV1().ConfigMaps("").Update(ctx, cmK8s, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	} else {
		_, err = global.KubeConfigSet.CoreV1().ConfigMaps("").Create(ctx, cmK8s, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
