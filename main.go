package main

import (
	"k8sdashboar.com/global"
	"k8sdashboar.com/initiallize"
)

func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8S()
	//println(global.CONF.System.Addr)
	panic(r.Run(global.CONF.System.Addr))
}
