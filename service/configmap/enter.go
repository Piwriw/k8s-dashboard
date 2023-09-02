package configmap

import "k8sdashboar.com/convert"

type ConfigMapGroup struct {
	ConfigMapService
}

var cmConvert = convert.ConvertGroupApp.CMConvert
