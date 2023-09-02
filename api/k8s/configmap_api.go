package k8s

import (
	"github.com/gin-gonic/gin"
	configmap_req "k8sdashboar.com/model/configmap/request"
	"k8sdashboar.com/response"
)

type ConfigMapApi struct {
}

func (*ConfigMapApi) GetConfigMapDetailOrListHandler(c *gin.Context) {
	name := c.Query("name")
	keyword := c.Query("keyword")
	namespace := c.Param("namespace")
	if name == "" {
		list, err := configMapService.GetConfigMapList(namespace, keyword)
		if err != nil {
			response.FailWithMessage(c, "查询ConfigMap列表失败")
			return
		}
		response.SuccessWithDetailed(c, "查询ConfigMap列表成功", list)
		return
	} else {
		detail, err := configMapService.GetConfigMapDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, "查询ConfigMap详情失败")
			return
		}
		response.SuccessWithDetailed(c, "查询ConfigMap详情成功", detail)
		return
	}
}
func (*ConfigMapApi) CreateOrUpdateConfigMapHandler(c *gin.Context) {
	var configMapReq configmap_req.ConfigMap
	err := c.ShouldBindJSON(&configMapReq)
	if err != nil {
		response.FailWithMessage(c, "ConfigMap参数解析失败")
		return
	}
	err = configMapService.CreateOrUpdateConfigMap(configMapReq)
	if err != nil {
		response.FailWithMessage(c, "ConfigMap创建失败")
		return
	}
	response.Success(c)
}
func (*ConfigMapApi) DeleteConfigMapHandler(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := configMapService.DeleteConfigMap(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "ConfigMap删除失败")
		return
	}
	response.Success(c)
}
