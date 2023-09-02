package k8s

import (
	"github.com/gin-gonic/gin"
	pod_req "k8sdashboar.com/model/pod/request"
	"k8sdashboar.com/response"
)

type PodApi struct {
}

func (*PodApi) DeletePodHandler(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := podService.DeletePod(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "删除Pod失败"+err.Error())
		return
	}
	response.Success(c)
}
func (*PodApi) CreateOrUpdatePodHandler(c *gin.Context) {
	var podReq pod_req.Pod

	if err := c.ShouldBind(&podReq); err != nil {
		response.FailWithMessage(c, "参数解析失败："+err.Error())
		return
	}
	if err := podValidate.Validate(&podReq); err != nil {
		response.FailWithMessage(c, "参数验证失败："+err.Error())
		return
	}
	msg, err := podService.CreateOrUpdate(podReq)
	if err != nil {
		response.FailWithMessage(c, msg)
	}
	response.SuccessWithMessage(c, msg)
}
func (*PodApi) GetPodListOrDetailHandler(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		detail, err := podService.GetPodDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Pod详情成功", detail)
	} else {
		err, items := podService.GetPodList(namespace, keyword, c.Query("nodeName"))
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Pod列表成功", items)
	}
}
