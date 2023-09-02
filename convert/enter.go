package convert

import (
	"k8sdashboar.com/convert/configmap"
	"k8sdashboar.com/convert/node"
	"k8sdashboar.com/convert/pod"
	"k8sdashboar.com/convert/pv"
	"k8sdashboar.com/convert/pvc"
	"k8sdashboar.com/convert/sc"
	"k8sdashboar.com/convert/secret"
)

type ConvertGroup struct {
	PodConvert    pod.PodConvertGroup
	NodeConvert   node.NodeConvertGroup
	CMConvert     configmap.CMConvertGroup
	SecretConvert secret.SecretConvert
	PVConvert     pv.PVConvert
	PVCConvert    pvc.PVCConvert
	SCConvert     sc.SCConvert
}

var ConvertGroupApp = new(ConvertGroup)
