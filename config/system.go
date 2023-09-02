package config

type System struct {
	Addr          string `json:"addr" yaml:"addr"`
	Provisioner   string `json:"provisioner" yaml:"provisioner"`
	K8sConfigPath string `json:"k8SConfigPath" yaml:"k8SConfigPath"`
}
