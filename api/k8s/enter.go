package k8s

import (
	"k8sdashboar.com/service"
	"k8sdashboar.com/validate"
)

type ApiGroup struct {
	PodApi
	NameSpaceApi
	NodeApi
	ConfigMapApi
	SecretApi
	PVApi
	PVCApi
	SCApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup
var secretService = service.ServiceGroupApp.SecretServiceGroup
var pvService = service.ServiceGroupApp.PVServiceGroup
var pvcService = service.ServiceGroupApp.PVCServiceGroup
var scService = service.ServiceGroupApp.SCServiceGroup
