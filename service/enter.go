package service

import (
	"k8sdashboar.com/service/configmap"
	"k8sdashboar.com/service/node"
	"k8sdashboar.com/service/pod"
	"k8sdashboar.com/service/pv"
	"k8sdashboar.com/service/pvc"
	"k8sdashboar.com/service/sc"
	"k8sdashboar.com/service/secret"
)

type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.NodeServiceGroup
	ConfigMapServiceGroup configmap.ConfigMapGroup
	SecretServiceGroup    secret.SecretService
	PVServiceGroup        pv.PVService
	PVCServiceGroup       pvc.PVCGroup
	SCServiceGroup        sc.SCGroup
}

var ServiceGroupApp = new(ServiceGroup)
