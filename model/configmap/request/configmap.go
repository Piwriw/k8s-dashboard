package request

import "k8sdashboar.com/model/base"

type ConfigMap struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	Data      []base.ListMapItem `json:"data"`
}
