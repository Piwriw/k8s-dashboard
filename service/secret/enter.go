package secret

import "k8sdashboar.com/convert"

type SecretGroup struct {
	SecretService
}

var secretConvert = convert.ConvertGroupApp.SecretConvert
