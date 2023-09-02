package pod

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8sdashboar.com/model/base"
	pod_req "k8sdashboar.com/model/pod/request"
	"strconv"
	"strings"
)

const (
	probe_http = "http"
	probe_tcp  = "tcp"
	probe_exec = "exec"
)
const (
	volume_emptyDir  = "emptyDir"
	volume_configmap = "configmap"
	volume_secret    = "secret"
	volume_hostPath  = "hostPath"
	volume_downward  = "downwardAPI"
	volume_pvc       = "pvc"
)
const (
	SCHEDULING_NODENAME     = "nodeName"
	SCHEDULING_NODESELECTOR = "nodeSelector"
	SCHEDULING_NODEAFFINITY = "nodeAffinity"
	SCHEDULING_NODEANY      = "nodeAny"
)
const (
	ref_type_configmap = "configmap"
	ref_type_secret    = "secret"
)

type Req2K8sConvert struct {
}

// PodReq2K8s 把pod msg 转化为k8s
func (pc *Req2K8sConvert) PodReq2K8s(podReq pod_req.Pod) *corev1.Pod {
	nodeAffinity, nodeSelector, nodeName := pc.getK8sNodeScheduling(podReq)
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podReq.Base.Name,
			Namespace: podReq.Base.Namespace,
			Labels:    pc.getK8sLabels(podReq.Base.Labels),
		},
		Spec: corev1.PodSpec{
			NodeName:       nodeName,
			NodeSelector:   nodeSelector,
			Affinity:       nodeAffinity,
			Tolerations:    podReq.Tolerations,
			InitContainers: pc.getK8sContainers(podReq.InitContainers),
			Containers:     pc.getK8sContainers(podReq.Containers),
			Volumes:        pc.getK8sVolumes(podReq.Volumes),
			DNSConfig: &corev1.PodDNSConfig{
				Nameservers: podReq.NetWorking.DnsConfig.Nameserver,
			},
			DNSPolicy:     corev1.DNSPolicy(podReq.NetWorking.DnsPolicy),
			HostAliases:   pc.getK8sHostAliases(podReq.NetWorking.HostAliases),
			Hostname:      podReq.NetWorking.HostName,
			RestartPolicy: corev1.RestartPolicy(podReq.Base.RestartPolicy),
		},
	}
}
func (pc *Req2K8sConvert) getK8sNodeScheduling(podReq pod_req.Pod) (affinity *corev1.Affinity, nodeSelector map[string]string, nodeName string) {
	nodeScheduling := podReq.NodeScheduling
	switch nodeScheduling.Type {
	case SCHEDULING_NODENAME:
		nodeName = nodeScheduling.NodeName
		return
	case SCHEDULING_NODESELECTOR:
		nodeSelectorMap := make(map[string]string)
		for _, item := range nodeScheduling.NodeSelector {
			nodeSelectorMap[item.Key] = item.Value
		}
		nodeSelector = nodeSelectorMap
	case SCHEDULING_NODEAFFINITY:
		expressions := nodeScheduling.NodeAffinity
		matchExpression := make([]corev1.NodeSelectorRequirement, 0)

		for _, expression := range expressions {
			matchExpression = append(matchExpression, corev1.NodeSelectorRequirement{
				Key:      expression.Key,
				Operator: expression.Operator,
				Values:   strings.Split(expression.Value, ","),
			})
		}
		affinity = &corev1.Affinity{
			NodeAffinity: &corev1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{MatchExpressions: matchExpression},
					},
				},
				PreferredDuringSchedulingIgnoredDuringExecution: nil,
			},
		}
		return
	case SCHEDULING_NODEANY:
		//DO NOTing
	}
	return
}
func (pc *Req2K8sConvert) getK8sHostAliases(podReqHostAliases []base.ListMapItem) []corev1.HostAlias {
	podK8sHostAliases := make([]corev1.HostAlias, 0)
	for _, item := range podReqHostAliases {
		podK8sHostAliases = append(podK8sHostAliases, corev1.HostAlias{
			IP:        item.Key,
			Hostnames: strings.Split(item.Value, ","),
		})
	}
	return podK8sHostAliases
}
func (pc *Req2K8sConvert) getK8sVolumes(podReqVolumes []pod_req.Volume) []corev1.Volume {
	podK8sVolumes := make([]corev1.Volume, 0)
	for _, volume := range podReqVolumes {
		source := corev1.VolumeSource{}
		switch volume.Type {
		case volume_emptyDir:
			source = corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			}
		case volume_hostPath:
			pathType := volume.HostPathVolume.Type
			source = corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Type: &pathType,
					Path: volume.HostPathVolume.Path,
				},
			}
		case volume_configmap:
			source = corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: volume.ConfigMapRefVolume.Name,
				},
				Optional: &volume.ConfigMapRefVolume.Optional,
			}}
		case volume_secret:
			source = corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: volume.SecretRefVolume.Name,
					Optional:   &volume.ConfigMapRefVolume.Optional,
				},
			}
		case volume_downward:
			items := make([]corev1.DownwardAPIVolumeFile, 0)
			for _, item := range volume.DownwardApiVolume.Items {
				items = append(items, corev1.DownwardAPIVolumeFile{
					//容器内访问路径
					Path: item.Path,
					FieldRef: &corev1.ObjectFieldSelector{
						FieldPath: item.FiledRefPath,
					},
				})
			}
			source = corev1.VolumeSource{
				DownwardAPI: &corev1.DownwardAPIVolumeSource{
					Items: items,
				},
			}
		case volume_pvc:
			source = corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: volume.PVCVolume.Name,
			}}

		default:
			continue
		}

		podK8sVolumes = append(podK8sVolumes, corev1.Volume{
			Name:         volume.Name,
			VolumeSource: source,
		})
	}
	return podK8sVolumes
}
func (pc *Req2K8sConvert) getK8sContainers(podContainers []pod_req.Container) []corev1.Container {
	podK8sContainers := make([]corev1.Container, 0)
	for _, container := range podContainers {
		podK8sContainers = append(podK8sContainers, pc.getK8sContainer(container))
	}
	return podK8sContainers
}

func (pc *Req2K8sConvert) getK8sContainer(podReqContainer pod_req.Container) corev1.Container {
	return corev1.Container{
		Name:            podReqContainer.Name,
		Image:           podReqContainer.Image,
		ImagePullPolicy: corev1.PullPolicy(podReqContainer.ImagePullPolicy),
		TTY:             podReqContainer.Tty,
		Command:         podReqContainer.Command,
		Args:            podReqContainer.Args,
		WorkingDir:      podReqContainer.WorkingDir,
		SecurityContext: &corev1.SecurityContext{
			Privileged: &podReqContainer.Privileged,
		},
		Ports:          pc.getK8sPorts(podReqContainer.Ports),
		Env:            pc.getK8sEnv(podReqContainer.Envs),
		EnvFrom:        pc.getK8sEnvFrom(podReqContainer.EnvFrom),
		VolumeMounts:   pc.getK8sVolumeMounts(podReqContainer.VolumeMounts),
		StartupProbe:   pc.getK8sContainerProbe(podReqContainer.StartupProbe),
		LivenessProbe:  pc.getK8sContainerProbe(podReqContainer.LivenessProbe),
		ReadinessProbe: pc.getK8sContainerProbe(podReqContainer.ReadinessProbe),
		Resources:      pc.getK8sResources(podReqContainer.Resources),
	}
}
func (*Req2K8sConvert) getK8sEnvFrom(envFromReq []pod_req.EnvFromResource) []corev1.EnvFromSource {
	podK8sEnvFrom := make([]corev1.EnvFromSource, 0)
	for _, fromResource := range envFromReq {
		// 前缀
		envFromResource := corev1.EnvFromSource{
			Prefix: fromResource.Prefix,
		}
		switch fromResource.RefType {
		case ref_type_configmap:
			envFromResource.ConfigMapRef = &corev1.ConfigMapEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			break
		case ref_type_secret:
			envFromResource.SecretRef = &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			break
		}
	}
	return podK8sEnvFrom
}
func (*Req2K8sConvert) getK8sPorts(podReqContainerPorts []pod_req.ContainerPort) []corev1.ContainerPort {
	podK8sContainersPorts := make([]corev1.ContainerPort, 0)
	for _, item := range podReqContainerPorts {
		podK8sContainersPorts = append(podK8sContainersPorts, corev1.ContainerPort{
			Name:          item.Name,
			HostPort:      item.HostPort,
			ContainerPort: item.ContainerPort,
		})
	}
	return podK8sContainersPorts
}
func (*Req2K8sConvert) getK8sResources(podReqResource pod_req.Resources) corev1.ResourceRequirements {
	var k8sPodResource corev1.ResourceRequirements
	if !podReqResource.Enable {
		return k8sPodResource
	}
	k8sPodResource.Requests = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResource.CpuRequest)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResource.MemRequest)) + "Mi"),
	}
	k8sPodResource.Limits = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResource.CpuLimit)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResource.MemLimit)) + "Mi"),
	}
	return k8sPodResource
}
func (*Req2K8sConvert) getK8sContainerProbe(podReqProbe pod_req.ContainerProbe) *corev1.Probe {
	if podReqProbe.Enable {
		return nil
	}
	var k8sProbe corev1.Probe
	switch podReqProbe.Type {
	case probe_http:
		httpGet := podReqProbe.HttpGet
		k8sHttpHeaders := make([]corev1.HTTPHeader, 0)
		for _, header := range httpGet.HttpHeaders {
			k8sHttpHeaders = append(k8sHttpHeaders, corev1.HTTPHeader{
				Name:  header.Key,
				Value: header.Value,
			})
		}
		k8sProbe.HTTPGet = &corev1.HTTPGetAction{
			Scheme:      corev1.URIScheme(httpGet.Scheme),
			Host:        httpGet.Host,
			Port:        intstr.FromInt(int(httpGet.Port)),
			Path:        httpGet.Path,
			HTTPHeaders: k8sHttpHeaders,
		}
	case probe_tcp:
		tcpSocket := podReqProbe.TcpSocket
		k8sProbe.TCPSocket = &corev1.TCPSocketAction{
			Host: tcpSocket.Host,
			Port: intstr.FromInt(int(tcpSocket.Port)),
		}
	case probe_exec:
		exec := podReqProbe.Exec
		k8sProbe.Exec = &corev1.ExecAction{
			Command: exec.Command,
		}
	}
	return &k8sProbe
}
func (*Req2K8sConvert) getK8sVolumeMounts(podReqMounts []pod_req.VolumeMounts) []corev1.VolumeMount {
	podK8sVolumeMounts := make([]corev1.VolumeMount, 0)
	for _, mount := range podReqMounts {
		podK8sVolumeMounts = append(podK8sVolumeMounts, corev1.VolumeMount{
			Name:      mount.MountName,
			MountPath: mount.MountName,
			ReadOnly:  mount.ReadOnly,
		})
	}
	return podK8sVolumeMounts
}
func (*Req2K8sConvert) getK8sEnv(podReqEnv []pod_req.EnvVar) []corev1.EnvVar {
	podK8sEnvs := make([]corev1.EnvVar, 0)
	for _, item := range podReqEnv {
		envVar := corev1.EnvVar{
			Name: item.Name,
		}
		switch item.Type {
		case ref_type_configmap:
			envVar.ValueFrom = &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					Key: item.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: item.Name,
					},
				},
			}
			break
		case ref_type_secret:
			envVar.ValueFrom = &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: item.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: item.Name,
					},
				},
			}
			break
		default:
			envVar.Value = item.Value
		}

		podK8sEnvs = append(podK8sEnvs, envVar)
	}
	return podK8sEnvs
}

// getK8sLabels 转化为k8s的 labels
func (*Req2K8sConvert) getK8sLabels(podReqLabels []base.ListMapItem) map[string]string {
	podK8sLabels := make(map[string]string)
	for _, label := range podReqLabels {
		podK8sLabels[label.Key] = label.Value
	}
	return podK8sLabels
}
