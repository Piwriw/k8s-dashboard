package initiallize

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8sdashboar.com/global"
)

func K8S() {
	kubeconifg := global.CONF.System.K8sConfigPath
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconifg)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	global.KubeConfigSet = clientset
	if err != nil {
		panic(err.Error())
	}
}
