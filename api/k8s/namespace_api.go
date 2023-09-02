package k8s

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8sdashboar.com/global"
	namespace_res "k8sdashboar.com/model/namespace/response"
	"k8sdashboar.com/response"
)

type NameSpaceApi struct {
}

func (*NameSpaceApi) GetNamespaceListHandler(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	namespaceList := make([]namespace_res.Namespace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, namespace_res.Namespace{
			Name:            item.Name,
			CreateTimestamp: item.CreationTimestamp.Unix(),
			Status:          string(item.Status.Phase),
		})
	}
	response.SuccessWithDetailed(c, "success", namespaceList)
}
