package k8s

import (
	"github.com/gin-gonic/gin"
	pv_req "k8sdashboar.com/model/pv/request"
	"k8sdashboar.com/response"
)

type PVApi struct {
}

func (*PVApi) CreatePVHandler(c *gin.Context) {
	var pvReq pv_req.PersistentVolume
	err := c.ShouldBindJSON(&pvReq)
	if err != nil {
		response.FailWithMessage(c, "参数解析失败："+err.Error())
		return
	}

	err = pvService.CreatePV(pvReq)
	if err != nil {
		response.FailWithMessage(c, "创建或者更新PV失败")
		return
	}
	response.Success(c)
}
func (*PVApi) GetPVDetailOrListHandler(c *gin.Context) {
	name := c.Query("name")
	//namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	var data any
	var err error
	if name != "" {
		data, err = pvService.GetPVDetail(name)
	} else {
		data, err = pvService.GetPVList(keyword)
	}
	if err != nil {
		response.FailWithMessage(c, "获取PV失败，Detail:"+err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取PV成功", data)
}
func (*PVApi) DeletePVHandler(c *gin.Context) {
	name := c.Param("name")
	err := pvService.DeletePV(name)
	if err != nil {
		response.FailWithMessage(c, "删除Secret失败，Detail:"+err.Error())
		return
	}
	response.Success(c)
}
