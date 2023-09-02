package secret

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	global "k8sdashboar.com/global"
	secret_req "k8sdashboar.com/model/secret/request"
	secret_res "k8sdashboar.com/model/secret/response"
	"strings"
)

type SecretService struct {
}

func (*SecretService) GetSecretList(namespace, keyword string) ([]secret_res.Secret, error) {
	list, err := global.KubeConfigSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	secretResList := make([]secret_res.Secret, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		secretResList = append(secretResList, secretConvert.SecretK8s2ResItem(item))
	}
	return secretResList, nil
}

// GetSecretDetail 获取Secret详情
func (*SecretService) GetSecretDetail(namespace, name string) (*secret_res.Secret, error) {
	secretK8s, err := global.KubeConfigSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secretConvert.SecretK8s2ResDetail(*secretK8s)
	return &secretRes, err
}
func (*SecretService) DeleteSecret(namespace, name string) error {
	return global.KubeConfigSet.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

// CreateOrUpdateSecret 创建 或者 更新Secret
func (*SecretService) CreateOrUpdateSecret(secretReq secret_req.Secret) (secret corev1.Secret, err error) {
	secretK8s := secretConvert.SecretReq2K8sConvert(secretReq)
	// 查询是否存在
	ctx := context.TODO()
	serviceApi := global.KubeConfigSet.CoreV1().Secrets(secretReq.Namespace)
	_, errGet := serviceApi.Get(ctx, secretReq.Name, metav1.GetOptions{})
	if errGet == nil {
		_, err = serviceApi.Update(ctx, &secretK8s, metav1.UpdateOptions{})
	} else {
		_, err = serviceApi.Create(ctx, &secretK8s, metav1.CreateOptions{})
	}
	return
}
