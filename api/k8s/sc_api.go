package k8s

import (
	"github.com/gin-gonic/gin"
	sc_req "k8sdashboar.com/model/sc/request"
	"k8sdashboar.com/response"
)

type SCApi struct {
}

func (*SCApi) CreateSCHandler(c *gin.Context) {
	var scReq sc_req.StorageClass
	err := c.ShouldBindJSON(&scReq)
	if err != nil {
		response.FailWithMessage(c, "参数解析失败："+err.Error())
		return
	}

	err = scService.CreateSC(scReq)
	if err != nil {
		response.FailWithMessage(c, "创建或者更新StorageClass失败")
		return
	}
	response.Success(c)
}
func (*SCApi) GetSCDetailOrListHandler(c *gin.Context) {
	name := c.Query("name")
	//namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	var data any
	var err error
	if name != "" {
		data, err = scService.GetSCDetail(name)
	} else {
		data, err = scService.GetSCList(keyword)
	}
	if err != nil {
		response.FailWithMessage(c, "获取PV失败，Detail:"+err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取PV成功", data)
}
func (*SCApi) DeleteSCHandler(c *gin.Context) {
	name := c.Param("name")
	err := scService.DeleteSC(name)
	if err != nil {
		response.FailWithMessage(c, "删除SC失败，Detail:"+err.Error())
		return
	}
	response.Success(c)
}
