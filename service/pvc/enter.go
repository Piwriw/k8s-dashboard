package pvc

import "k8sdashboar.com/convert"

type PVCGroup struct {
	PVCService
}

var pvcConvert = convert.ConvertGroupApp.PVCConvert
