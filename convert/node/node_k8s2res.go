package node

import (
	corev1 "k8s.io/api/core/v1"
	"k8sdashboar.com/model/base"
	node_res "k8sdashboar.com/model/node/response"
)

type NodeK8s2Res struct {
}

func mapToList(m map[string]string) []base.ListMapItem {
	res := make([]base.ListMapItem, 0)
	for k, v := range m {
		res = append(res, base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return res
}
func (nk *NodeK8s2Res) GetNodeDetailRes(nodeK8s corev1.Node) node_res.Node {
	nodeRes := nk.GetNodeItemRes(nodeK8s)
	nodeRes.Taints = nodeK8s.Spec.Taints
	nodeRes.Labels = mapToList(nodeK8s.Labels)
	return nodeRes
}
func (nk *NodeK8s2Res) GetNodeItemRes(nodeK8s corev1.Node) node_res.Node {
	return node_res.Node{
		Name:             nodeK8s.Name,
		Status:           nk.getNodeStatusRes(nodeK8s.Status.Conditions),
		Age:              nodeK8s.CreationTimestamp.Unix(),
		InternalIP:       nk.getNodeIP(nodeK8s.Status.Addresses, corev1.NodeInternalIP),
		ExternalIP:       nk.getNodeIP(nodeK8s.Status.Addresses, corev1.NodeExternalIP),
		Version:          nodeK8s.Status.NodeInfo.KubeletVersion,
		OsImage:          nodeK8s.Status.NodeInfo.OSImage,
		KernelVersion:    nodeK8s.Status.NodeInfo.KernelVersion,
		ContainerRuntime: nodeK8s.Status.NodeInfo.ContainerRuntimeVersion,
	}
}
func (*NodeK8s2Res) getNodeIP(address []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
	for _, item := range address {
		if item.Type == addressType {
			return item.Address
		}
	}
	return "<none>"
}
func (*NodeK8s2Res) getNodeStatusRes(nodeConditions []corev1.NodeCondition) string {
	nodeStatus := "NotReady"
	for _, condition := range nodeConditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			nodeStatus = "Ready"
			break
		}
	}
	return nodeStatus
}
