package k8s

import (
	"github.com/gin-gonic/gin"
	pvc_req "k8sdashboar.com/model/pvc/request"
	"k8sdashboar.com/response"
)

type PVCApi struct {
}

func (*PVCApi) CreatePVCHandler(c *gin.Context) {
	var pvcReq pvc_req.PersistentVolumeClaim
	err := c.ShouldBindJSON(&pvcReq)
	if err != nil {
		response.FailWithMessage(c, "参数解析失败："+err.Error())
		return
	}

	err = pvcService.CreatePVC(pvcReq)
	if err != nil {
		response.FailWithMessage(c, "创建或者更新PVC失败")
		return
	}
	response.Success(c)
}
func (*PVCApi) GetPVCDetailOrListHandler(c *gin.Context) {
	name := c.Query("name")
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	var data any
	var err error
	if name != "" {
		data, err = pvcService.GetPVCDetail(name, namespace)
	} else {
		data, err = pvcService.GetPVCList(keyword, namespace)
	}
	if err != nil {
		response.FailWithMessage(c, "获取PVC失败，Detail:"+err.Error())
		return
	}
	response.SuccessWithDetailed(c, "获取PVC成功", data)
}
func (*PVCApi) DeletePVCHandler(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	err := pvcService.DeletePVC(name, namespace)
	if err != nil {
		response.FailWithMessage(c, "删除Secret失败，Detail:"+err.Error())
		return
	}
	response.Success(c)
}
