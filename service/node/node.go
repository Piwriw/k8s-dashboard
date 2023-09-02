package node

import (
	"context"
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8sdashboar.com/global"
	node_req "k8sdashboar.com/model/node/request"
	node_res "k8sdashboar.com/model/node/response"
	"strings"
)

type NodeService struct {
}

func (*NodeService) UpdateNodeTaint(updateTaint node_req.UpdateTaint) error {
	patchData := map[string]any{
		"spec": map[string]any{
			"taints": updateTaint.Taints,
		},
	}
	patchBytes, _ := json.Marshal(&patchData)
	ctx := context.TODO()
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(ctx,
		updateTaint.Name,
		types.StrategicMergePatchType,
		patchBytes, metav1.PatchOptions{})
	return err
}

// UpdateNodeLabel 更新节点标签
func (*NodeService) UpdateNodeLabel(updateLabel node_req.UpdateLabel) error {
	labelsMap := make(map[string]string, 0)
	for _, label := range updateLabel.Labels {
		labelsMap[label.Key] = label.Value
	}
	// 这样才会替换
	labelsMap["$patch"] = "replace"

	patchData := map[string]any{
		"metadata": map[string]any{
			"labels": labelsMap,
		},
	}
	patchBytes, _ := json.Marshal(&patchData)
	ctx := context.TODO()
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(ctx,
		updateLabel.Name,
		types.StrategicMergePatchType,
		patchBytes, metav1.PatchOptions{})
	return err
}

// GetNodeDetail 获取 Node 详情
func (*NodeService) GetNodeDetail(nodeName string) (*node_res.Node, error) {
	ctx := context.TODO()
	nodeK8s, err := global.KubeConfigSet.CoreV1().Nodes().Get(ctx, nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	detailRes := nodeConvert.GetNodeDetailRes(*nodeK8s)
	return &detailRes, nil
}

// GetNodeList 获取 Node 列表
func (*NodeService) GetNodeList(keyword string) ([]node_res.Node, error) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodeResList := make([]node_res.Node, 0)
	for _, item := range list.Items {
		if strings.Contains(item.Name, keyword) {
			nodeRes := nodeConvert.GetNodeItemRes(item)
			nodeResList = append(nodeResList, nodeRes)
		}
	}
	return nodeResList, err
}
