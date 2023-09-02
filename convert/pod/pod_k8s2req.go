package pod

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	pod_req "k8sdashboar.com/model/pod/request"
	"strings"
)

const volume_type_emptydir = "emptyDir"

type K8s2rReqConvert struct {
	volumeMap map[string]string
}

func (kr *K8s2rReqConvert) PodK8s2Req(podK8s corev1.Pod) pod_req.Pod {
	return pod_req.Pod{
		Base:           kr.getReqBase(podK8s),
		Tolerations:    podK8s.Spec.Tolerations,
		NodeScheduling: kr.getReqNodeScheduling(podK8s),
		Volumes:        kr.getReqVolume(podK8s.Spec.Volumes),
		NetWorking:     kr.getReqNetworking(podK8s),
	}
}
func (kr *K8s2rReqConvert) getReqNodeScheduling(podK8s corev1.Pod) pod_req.NodeScheduling {
	nodeScheduling := pod_req.NodeScheduling{
		Type: SCHEDULING_NODEANY,
	}
	if podK8s.Spec.NodeSelector != nil {
		nodeScheduling.Type = SCHEDULING_NODESELECTOR
		labels := make([]base.ListMapItem, 0)
		for k, v := range podK8s.Spec.NodeSelector {
			labels = append(labels, base.ListMapItem{
				Key:   k,
				Value: v,
			})
		}
		nodeScheduling.NodeSelector = labels
		return nodeScheduling
	}
	if podK8s.Spec.Affinity != nil {
		nodeScheduling.Type = SCHEDULING_NODEAFFINITY
		term := podK8s.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms[0]
		matchExpressions := make([]pod_req.NodeSelectorTermExpressions, 0)
		for _, expression := range term.MatchExpressions {
			matchExpressions = append(matchExpressions, pod_req.NodeSelectorTermExpressions{
				Key:      expression.Key,
				Operator: expression.Operator,
				Value:    strings.Join(expression.Values, ","),
			})
		}
		nodeScheduling.NodeAffinity = matchExpressions
		return nodeScheduling
	}
	if podK8s.Spec.NodeName != "" {
		nodeScheduling.Type = SCHEDULING_NODENAME
		nodeScheduling.NodeName = podK8s.Spec.NodeName
		return nodeScheduling
	}
	return nodeScheduling
}
func (kr *K8s2rReqConvert) getReqNetworking(podK8s corev1.Pod) pod_req.NetWorking {
	return pod_req.NetWorking{
		HostNetwork: podK8s.Spec.HostNetwork,
		HostName:    podK8s.Spec.Hostname,
		DnsPolicy:   string(podK8s.Spec.DNSPolicy),
		DnsConfig:   kr.getReqDnsConfig(podK8s.Spec.DNSConfig),
		HostAliases: kr.getReqHostAliases(podK8s.Spec.HostAliases),
	}
}
func (kr *K8s2rReqConvert) getReqContainers(podContainersK8s []corev1.Container) {
	podReqContainers := make([]pod_req.Container, 0)
	for _, item := range podContainersK8s {
		podReqContainers = append(podReqContainers, kr.getReqContainer(item))
	}
}
func (kr *K8s2rReqConvert) getReqContainer(containerK8s corev1.Container) pod_req.Container {
	return pod_req.Container{
		Name:            containerK8s.Name,
		Image:           containerK8s.Image,
		ImagePullPolicy: string(containerK8s.ImagePullPolicy),
		Tty:             containerK8s.TTY,
		Ports:           kr.getReqContainerPorts(containerK8s.Ports),
		WorkingDir:      containerK8s.WorkingDir,
		Command:         containerK8s.Command,
		Args:            containerK8s.Args,
		Envs:            kr.getReqContainerEnvs(containerK8s.Env),
		EnvFrom:         kr.getReqContainerEnvFrom(containerK8s.EnvFrom),
		Privileged:      kr.getReqContainerPrivileged(containerK8s.SecurityContext),
		Resources:       kr.getReqContainerResource(containerK8s.Resources),
		VolumeMounts:    kr.getReqContainerVolumeMounts(containerK8s.VolumeMounts),
		StartupProbe:    kr.getReqContainerProbe(containerK8s.StartupProbe),
		LivenessProbe:   kr.getReqContainerProbe(containerK8s.LivenessProbe),
		ReadinessProbe:  kr.getReqContainerProbe(containerK8s.ReadinessProbe),
	}
}
func (kr *K8s2rReqConvert) getReqContainerEnvFrom(k8sEnvFromList []corev1.EnvFromSource) []pod_req.EnvFromResource {
	podEnvFromList := make([]pod_req.EnvFromResource, 0)
	for _, envFromItem := range k8sEnvFromList {
		envFromReq := pod_req.EnvFromResource{
			Name: envFromItem.Prefix,
		}
		if envFromItem.ConfigMapRef != nil {
			envFromReq.Name = envFromItem.ConfigMapRef.Name
			envFromReq.RefType = ref_type_configmap
		}
		if envFromItem.SecretRef != nil {
			envFromReq.Name = envFromItem.SecretRef.Name
			envFromReq.RefType = ref_type_secret
		}
	}
	return podEnvFromList
}
func (kr *K8s2rReqConvert) getReqContainerProbe(probeK8s *corev1.Probe) pod_req.ContainerProbe {
	containerProbe := pod_req.ContainerProbe{
		Enable: false,
	}
	//先判断是否探针为空
	if probeK8s != nil {
		containerProbe.Enable = true
		//再判断 探针具体是什么类型
		if probeK8s.Exec != nil {
			containerProbe.Type = probe_exec
			containerProbe.Exec.Command = probeK8s.Exec.Command
		} else if probeK8s.HTTPGet != nil {
			containerProbe.Type = probe_http
			httpGet := probeK8s.HTTPGet
			headersReq := make([]base.ListMapItem, 0)
			for _, headerK8s := range httpGet.HTTPHeaders {
				headersReq = append(headersReq, base.ListMapItem{
					Key:   headerK8s.Name,
					Value: headerK8s.Value,
				})
			}
			containerProbe.HttpGet = pod_req.ProbeHttpGet{
				Host:        httpGet.Host,
				Port:        httpGet.Port.IntVal,
				Scheme:      string(httpGet.Scheme),
				Path:        httpGet.Path,
				HttpHeaders: headersReq,
			}
		} else if probeK8s.TCPSocket != nil {
			containerProbe.Type = probe_tcp
			containerProbe.TcpSocket = pod_req.ProbeTcpSocket{
				Host: probeK8s.TCPSocket.Host,
				Port: probeK8s.TCPSocket.Port.IntVal,
			}
		} else {
			containerProbe.Type = probe_http
			return containerProbe
		}
		containerProbe.InitialDelaySeconds = probeK8s.InitialDelaySeconds
		containerProbe.PeriodSeconds = probeK8s.PeriodSeconds
		containerProbe.TimeoutSeconds = probeK8s.TimeoutSeconds
		containerProbe.SuccessThreshold = probeK8s.SuccessThreshold
		containerProbe.FailureThreshold = probeK8s.FailureThreshold
	}
	return containerProbe
}
func (kr *K8s2rReqConvert) getReqContainerVolumeMounts(volumeMountK8s []corev1.VolumeMount) []pod_req.VolumeMounts {
	reqVolumeMounts := make([]pod_req.VolumeMounts, 0)
	for _, item := range volumeMountK8s {
		//非emptydir 过滤掉
		_, ok := kr.volumeMap[item.Name]
		if ok {
			reqVolumeMounts = append(reqVolumeMounts, pod_req.VolumeMounts{
				MountName: item.Name,
				MountPath: item.MountPath,
				ReadOnly:  item.ReadOnly,
			})
		}
	}
	return reqVolumeMounts
}
func (kr *K8s2rReqConvert) getReqContainerResource(requirementsK8s corev1.ResourceRequirements) pod_req.Resources {
	reqResources := pod_req.Resources{
		Enable: false,
	}
	requests := requirementsK8s.Requests
	limits := requirementsK8s.Limits
	if requests != nil {
		reqResources.Enable = true
		reqResources.CpuRequest = int32(requests.Cpu().MilliValue()) // m
		//MiB
		reqResources.MemRequest = int32(requests.Memory().Value() / (1024 * 1024)) //Bytes
	}
	if limits != nil {
		reqResources.Enable = true
		reqResources.CpuLimit = int32(limits.Cpu().MilliValue())
		reqResources.MemLimit = int32(limits.Memory().Value() / (1024 * 1024))
	}
	return reqResources
}
func (kr *K8s2rReqConvert) getReqContainerPrivileged(ctx *corev1.SecurityContext) (privileged bool) {
	if ctx != nil {
		privileged = *ctx.Privileged
	}
	return
}
func (kr *K8s2rReqConvert) getReqContainerEnvs(envsK8s []corev1.EnvVar) []pod_req.EnvVar {
	envsReq := make([]pod_req.EnvVar, 0)
	for _, item := range envsK8s {
		envVar := pod_req.EnvVar{
			Name: item.Name,
		}
		if item.ValueFrom != nil {
			if item.ValueFrom.ConfigMapKeyRef != nil {
				envVar.Type = ref_type_configmap
				envVar.Value = item.ValueFrom.ConfigMapKeyRef.Key
			}
			if item.ValueFrom.SecretKeyRef != nil {
				envVar.Type = ref_type_secret
				envVar.Value = item.ValueFrom.SecretKeyRef.Key
				envVar.RefName = item.ValueFrom.SecretKeyRef.Name
			}
		} else {
			envVar.Value = item.Value
		}
		envsReq = append(envsReq, envVar)
	}
	return envsReq
}
func (kr *K8s2rReqConvert) getReqContainerPorts(portsK8s []corev1.ContainerPort) []pod_req.ContainerPort {
	portsReq := make([]pod_req.ContainerPort, 0)
	for _, item := range portsK8s {
		portsReq = append(portsReq, pod_req.ContainerPort{
			Name:          item.Name,
			ContainerPort: item.ContainerPort,
			HostPort:      item.HostPort,
		})
	}
	return portsReq
}
func (kr *K8s2rReqConvert) getReqVolume(volumes []corev1.Volume) []pod_req.Volume {
	volumesReq := make([]pod_req.Volume, 0)
	if kr.volumeMap == nil {
		kr.volumeMap = make(map[string]string)
	}
	for _, volume := range volumes {
		var volumeReq *pod_req.Volume
		if volume.EmptyDir != nil {
			volumeReq = &pod_req.Volume{
				Type: volume_emptyDir,
				Name: volume.Name,
			}
		}

		if volume.ConfigMap != nil {
			var optional bool
			if volume.ConfigMap.Optional != nil {
				optional = *volume.ConfigMap.Optional
			}
			volumeReq.Type = volume_configmap
			volumeReq = &pod_req.Volume{
				Type: volume_type_emptydir,
				Name: volume.Name,
				ConfigMapRefVolume: pod_req.ConfigMapRefVolume{
					Name:     volume.ConfigMap.Name,
					Optional: optional,
				},
			}
		}

		if volume.Secret != nil {
			var optional bool
			if volume.Secret.Optional != nil {
				optional = *volume.ConfigMap.Optional
			}
			volumeReq = &pod_req.Volume{
				Type: volume_secret,
				SecretRefVolume: pod_req.SecretRefVolume{
					Name:     volume.Secret.SecretName,
					Optional: optional,
				},
			}
		}

		if volume.HostPath != nil {
			volumeReq = &pod_req.Volume{
				Type: volume_hostPath,
				HostPathVolume: pod_req.HostPathVolume{
					Path: volume.HostPath.Path,
					Type: *volume.HostPath.Type,
				},
			}
		}

		if volume.PersistentVolumeClaim != nil {
			volumeReq = &pod_req.Volume{
				Type: volume_pvc,
				Name: volume.Name,
				PVCVolume: pod_req.PVCVolume{
					Name: volume.PersistentVolumeClaim.ClaimName,
				},
			}
		}
		if volume.DownwardAPI != nil {
			items := make([]pod_req.DownwardApiVolumeItem, 0)
			for _, item := range volume.DownwardAPI.Items {
				items = append(items, pod_req.DownwardApiVolumeItem{
					Path:         item.Path,
					FiledRefPath: item.FieldRef.FieldPath,
				})
			}
			volumeReq = &pod_req.Volume{
				Type: volume_downward,
				Name: volume.Name,
				DownwardApiVolume: pod_req.DownwardApiVolume{
					Items: items,
				},
			}
		}

		if volumeReq == nil {
			continue
		}

		kr.volumeMap[volume.Name] = ""
		volumesReq = append(volumesReq, *volumeReq)
	}
	return volumesReq
}
func (*K8s2rReqConvert) getReqDnsConfig(dnsConfigK8s *corev1.PodDNSConfig) pod_req.DnsConfig {
	var dnsConfigReq pod_req.DnsConfig
	if dnsConfigK8s != nil {
		dnsConfigReq.Nameserver = dnsConfigK8s.Nameservers
	}
	return dnsConfigReq
}
func (*K8s2rReqConvert) getReqHostAliases(hostAliases []corev1.HostAlias) []base.ListMapItem {
	hostAliasesReq := make([]base.ListMapItem, 0)
	for _, alias := range hostAliases {
		hostAliasesReq = append(hostAliasesReq, base.ListMapItem{
			Key:   alias.IP,
			Value: strings.Join(alias.Hostnames, ","),
		})
	}
	return hostAliasesReq
}
func (*K8s2rReqConvert) getReqLabels(data map[string]string) []base.ListMapItem {
	labels := make([]base.ListMapItem, 0)
	for k, v := range data {
		labels = append(labels, base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return labels
}
func (kr *K8s2rReqConvert) getReqBase(podK8s corev1.Pod) pod_req.Base {
	return pod_req.Base{
		Name:          podK8s.Name,
		Namespace:     podK8s.Namespace,
		Labels:        kr.getReqLabels(podK8s.Labels),
		RestartPolicy: string(podK8s.Spec.RestartPolicy),
	}
}
