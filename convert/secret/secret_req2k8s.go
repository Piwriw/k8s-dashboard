package secret

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/model/base"
	secret_req "k8sdashboar.com/model/secret/request"
)

type SecretReq2K8s struct {
}

func (*SecretReq2K8s) SecretReq2K8sConvert(secretReq secret_req.Secret) corev1.Secret {
	return corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretReq.Name,
			Namespace: secretReq.Namespace,
			Labels:    base.ToMap(secretReq.Labels),
		},
		Type:       secretReq.Type,
		StringData: base.ToMap(secretReq.Data),
	}
}
