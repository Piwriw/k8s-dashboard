package router

import (
	"k8sdashboar.com/router/example"
	"k8sdashboar.com/router/k8s"
)

type RouterGroup struct {
	ExampleRouterApp example.ExampleRouter
	InitK8SRouterApp k8s.K8SRouter
}

var RouterGroupApp = new(RouterGroup)
