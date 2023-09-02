package pod

import "k8sdashboar.com/convert"

type PodServiceGroup struct {
	PodService
}

var podConvert = convert.ConvertGroupApp.PodConvert
