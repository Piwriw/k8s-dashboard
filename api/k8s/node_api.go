package k8s

import (
	"github.com/gin-gonic/gin"
	node_req "k8sdashboar.com/model/node/request"
	"k8sdashboar.com/response"
)

type NodeApi struct {
}

func (*NodeApi) UpdateNodeTaintHandler(c *gin.Context) {
	var updatedTaint node_req.UpdateTaint
	err := c.ShouldBind(&updatedTaint)
	if err != nil {
		response.FailWithMessage(c, "参数解析错误")
		return
	}
	err = nodeService.UpdateNodeTaint(updatedTaint)
	if err != nil {
		response.FailWithMessage(c, "更新节点污点（Taint）错误,detail:"+err.Error())
		return
	}
	response.Success(c)
}
func (*NodeApi) UpdateNodeLabelHandler(c *gin.Context) {
	var updatedLabel node_req.UpdateLabel
	err := c.ShouldBind(&updatedLabel)
	if err != nil {
		response.FailWithMessage(c, "参数解析错误")
		return
	}
	err = nodeService.UpdateNodeLabel(updatedLabel)
	if err != nil {
		response.FailWithMessage(c, "更新结点标签错误")
		return
	}
	response.Success(c)
}
func (*NodeApi) GetNodeDetailOrListHandler(c *gin.Context) {
	keyword := c.Query("keyword")
	nodeName := c.Query("nodeName")
	if nodeName != "" {
		list, err := nodeService.GetNodeDetail(nodeName)
		if err != nil {
			response.FailWithMessage(c, "获取Node详情失败")
			return
		}
		response.SuccessWithDetailed(c, "获取Node详情成功", list)
	} else {
		list, err := nodeService.GetNodeList(keyword)
		if err != nil {
			response.FailWithMessage(c, "获取Node列表失败,Detail:"+err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Node列表成功", list)
	}

}
