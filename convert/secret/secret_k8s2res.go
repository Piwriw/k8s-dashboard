package secret

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	secret_res "k8sdashboar.com/model/secret/response"
)

type SecretK8s2Res struct {
}

func (*SecretK8s2Res) SecretK8s2ResItem(secretK8s corev1.Secret) secret_res.Secret {
	return secret_res.Secret{
		Name:      secretK8s.Name,
		Namespace: secretK8s.Namespace,
		Type:      secretK8s.Type,
		DataNum:   len(secretK8s.StringData),
		Age:       secretK8s.CreationTimestamp.Unix(),
	}
}
func (*SecretK8s2Res) SecretK8s2ResDetail(secretK8s corev1.Secret) secret_res.Secret {
	return secret_res.Secret{
		Name:      secretK8s.Name,
		Namespace: secretK8s.Namespace,
		Type:      secretK8s.Type,
		Data:      base.ToListWithMapByte(secretK8s.Data),
		Labels:    base.ToList(secretK8s.Labels),
	}
}
