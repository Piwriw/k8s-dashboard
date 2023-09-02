package request

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
)

type Volume struct {
	Name string `json:"name"`
	// emptyDir | configMap | secret | hostPath | downward | pvc
	Type               string             `json:"type"`
	ConfigMapRefVolume ConfigMapRefVolume `json:"configMapRefVolume"`
	SecretRefVolume    SecretRefVolume    `json:"secretRefVolume"`
	HostPathVolume     HostPathVolume     `json:"hostPathVolume"`
	DownwardApiVolume  DownwardApiVolume  `json:"downwardApiVolume"`
	PVCVolume          PVCVolume          `json:"PVCVolume"`
}
type PVCVolume struct {
	Name string `json:"name"`
}
type DownwardApiVolumeItem struct {
	Path         string `json:"path"`
	FiledRefPath string `json:"filedRefPath"`
}
type DownwardApiVolume struct {
	Items []DownwardApiVolumeItem `json:"items"`
}
type HostPathVolume struct {
	Type corev1.HostPathType `json:"type"`
	// 宿主机路径
	Path string `json:"path"`
}
type ConfigMapRefVolume struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
}
type SecretRefVolume struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
}

// Resources 资源
type Resources struct {
	Enable bool `json:"enable"`
	//Mi
	MemRequest int32 `json:"memRequest"`
	MemLimit   int32 `json:"memLimit"`
	//m
	CpuRequest int32 `json:"cpuRequest"`
	CpuLimit   int32 `json:"cpuLimit"`
}

// VolumeMounts : 卷挂载
type VolumeMounts struct {
	MountName string `json:"mountName"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}
type EnvVar struct {
	Name    string `json:"name"`
	RefName string `json:"refName"`
	Value   string `json:"value"`
	// configmap |secret |default(k/v)
	Type string `json:"type"`
}
type EnvFromResource struct {
	// 资源名称
	Name    string `json:"name"`
	RefType string `json:"refType"`
	Prefix  string `json:"prefix"`
}

// Container 容器
type Container struct {
	Name            string            `json:"name"`
	Image           string            `json:"image"`
	ImagePullPolicy string            `json:"imagePullPolicy"`
	Tty             bool              `json:"tty"` //tty 字段表示是否需要给容器分配一个终端
	Ports           []ContainerPort   `json:"ports"`
	WorkingDir      string            `json:"workingDir"`
	Command         []string          `json:"command"`
	Args            []string          `json:"args"` //参数
	Envs            []EnvVar          `json:"envs"` //环境变量
	EnvFrom         []EnvFromResource `json:"envFrom"`
	Privileged      bool              `json:"privileged"` //特权模式
	Resources       Resources         `json:"resources"`
	VolumeMounts    []VolumeMounts    `json:"volumeMounts"`   // 卷挂载
	StartupProbe    ContainerProbe    `json:"startupProbe"`   //启动探针
	LivenessProbe   ContainerProbe    `json:"livenessProbe"`  //存活探针
	ReadinessProbe  ContainerProbe    `json:"readinessProbe"` //就绪探针
}
type ContainerPort struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}
type DnsConfig struct {
	Nameserver []string `json:"nameserver"`
}
type NetWorking struct {
	HostNetwork bool               `json:"hostNetwork"`
	HostName    string             `json:"hostName"`
	DnsPolicy   string             `json:"dnsPolicy"`
	DnsConfig   DnsConfig          `json:"dnsConfig"`
	HostAliases []base.ListMapItem `json:"hostAliases"`
}
type Base struct {
	Name          string             `json:"name"`
	Labels        []base.ListMapItem `json:"labels"`
	Namespace     string             `json:"namespace"`
	RestartPolicy string             `json:"restartPolicy"`
}
type ProbeHttpGet struct {
	Scheme      string             `json:"scheme"` //请求协议
	Host        string             `json:"host"`   // 请求host 如果为空，那就Pod内请求
	Path        string             `json:"path"`
	Port        int32              `json:"port"`
	HttpHeaders []base.ListMapItem `json:"httpHeaders"`
}
type ProbeCommand struct {
	Command []string `json:"command"`
}
type ProbeTime struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds"` //初始化时间，初始化诺干秒之后开始探针
	PeriodSeconds       int32 `json:"periodSeconds"`       //每隔诺干秒之后，开始探针
	TimeoutSeconds      int32 `json:"timeoutSeconds"`      //探针等待时间
	SuccessThreshold    int32 `json:"successThreshold"`    //探针若干次，才认为这次探针成功
	FailureThreshold    int32 `json:"failureThreshold"`    //探针若干次，才认为这次探针失败
}
type ProbeTcpSocket struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
type ContainerProbe struct {
	Enable    bool           `json:"enable"` //是否打开探针
	Type      string         `json:"type"`   //TCP/HTTP/EXEC
	HttpGet   ProbeHttpGet   `json:"httpGet"`
	Exec      ProbeCommand   `json:"exec"`
	TcpSocket ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}
type NodeSelectorTermExpressions struct {
	Key      string                      `json:"key"`
	Operator corev1.NodeSelectorOperator `json:"operator"`
	Value    string                      `json:"value"`
}
type NodeScheduling struct {
	Type         string                        `json:"type"`
	NodeName     string                        `json:"nodeName"`
	NodeSelector []base.ListMapItem            `json:"nodeSelector"`
	NodeAffinity []NodeSelectorTermExpressions `json:"nodeAffinity"`
}
type Pod struct {
	Base        Base                `json:"base"`
	Tolerations []corev1.Toleration `json:"tolerations"`
	// 调度策略
	NodeScheduling NodeScheduling `json:"nodeScheduling"`
	Volumes        []Volume       `json:"volumes"`
	NetWorking     NetWorking     `json:"netWorking"`
	InitContainers []Container    `json:"initContainers"`
	Containers     []Container    `json:"containers"`
}
