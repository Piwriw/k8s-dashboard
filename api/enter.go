package api

import (
	"k8sdashboar.com/api/example"
	"k8sdashboar.com/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ExampleApi
	K8SApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
