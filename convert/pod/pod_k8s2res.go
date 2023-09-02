package pod

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	pod_res "k8sdashboar.com/model/pod/response"
)

type K8s2rResConvert struct {
}

func (kr *K8s2rResConvert) PodK8s2ItemRes(podK8s corev1.Pod) pod_res.PodListItem {
	return pod_res.PodListItem{
		Name:     podK8s.Name,
		Ready:    kr.getResReady(podK8s),
		Status:   kr.getResStatus(podK8s),
		Restarts: kr.getResRestarts(podK8s),
		Age:      podK8s.CreationTimestamp.Unix(),
		IP:       podK8s.Status.PodIP,
		Node:     podK8s.Spec.NodeName,
	}
}
func (*K8s2rResConvert) getResRestarts(podK8s corev1.Pod) int32 {
	var restartC int32
	for _, containerStatus := range podK8s.Status.ContainerStatuses {
		restartC += containerStatus.RestartCount
	}
	return restartC
}
func (*K8s2rResConvert) getResStatus(podK8s corev1.Pod) string {
	// 不完全匹配
	var podStatus string
	if podK8s.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}
	return podStatus
}
func (*K8s2rResConvert) getResReady(podK8s corev1.Pod) string {
	var totalC, readyC int
	for _, containerStatus := range podK8s.Status.ContainerStatuses {
		if containerStatus.Ready {
			readyC++
		}
		totalC++
	}
	return fmt.Sprintf("%d/%d", readyC, totalC)
}
