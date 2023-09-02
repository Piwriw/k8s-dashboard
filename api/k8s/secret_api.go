package k8s

import (
	"github.com/gin-gonic/gin"
	secret_req "k8sdashboar.com/model/secret/request"
	"k8sdashboar.com/response"
)

type SecretApi struct {
}

func (*SecretApi) GetSecretDetailOrListHandler(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	var data any
	var err error
	if name != "" {
		data, err = secretService.GetSecretDetail(namespace, name)
	} else {
		data, err = secretService.GetSecretList(namespace, keyword)
	}
	if err != nil {
		response.FailWithMessage(c, "获取Secret失败，Detail:"+err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取Secret成功", data)
}
func (*SecretApi) DeleteSecretHandler(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := secretService.DeleteSecret(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "删除Secret失败，Detail:"+err.Error())
		return
	}
	response.Success(c)
}
func (*SecretApi) CreateOrUpdateSecretHandler(c *gin.Context) {
	var secretReq secret_req.Secret
	err := c.ShouldBindJSON(&secretReq)
	if err != nil {
		response.FailWithMessage(c, "参数解析失败："+err.Error())
		return
	}
	_, err = secretService.CreateOrUpdateSecret(secretReq)
	if err != nil {
		response.FailWithMessage(c, "创建或者更新Secret失败")
		return
	}
	response.Success(c)
}
