package k8s

import (
	"github.com/gin-gonic/gin"
	"k8sdashboar.com/api"
	"k8sdashboar.com/api/k8s"
)

type K8SRouter struct {
}

func (kr *K8SRouter) InitK8SRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	k8SApiGroup := api.ApiGroupApp.K8SApiGroup
	kr.initPodRouter(group, k8SApiGroup)
	kr.initNodeRouter(group, k8SApiGroup)
	kr.initConfigMapRouter(group, k8SApiGroup)
	kr.initPVRouter(group, k8SApiGroup)
	kr.initPVCRouter(group, k8SApiGroup)
	kr.initSCRouter(group, k8SApiGroup)
}
func (*K8SRouter) initSCRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/sc/:namespace", k8SApiGroup.GetSCDetailOrListHandler)
	group.DELETE("/sc/:namespace/:name", k8SApiGroup.DeleteSCHandler)
	group.POST("/sc", k8SApiGroup.CreateSCHandler)
}
func (*K8SRouter) initPVCRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/pvc/:namespace", k8SApiGroup.GetPVCDetailOrListHandler)
	group.DELETE("/pvc/:namespace/:name", k8SApiGroup.DeletePVCHandler)
	group.POST("/pvc", k8SApiGroup.CreatePVCHandler)
}

func (*K8SRouter) initPVRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/pv/:namespace", k8SApiGroup.GetPVDetailOrListHandler)
	group.DELETE("/pv/:namespace/:name", k8SApiGroup.DeletePVHandler)
	group.POST("/pv", k8SApiGroup.CreatePVHandler)
}

func (*K8SRouter) initSecretRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/secret/:namespace", k8SApiGroup.GetSecretDetailOrListHandler)
	group.DELETE("/secret/:namespace/:name", k8SApiGroup.DeleteSecretHandler)
	group.POST("/secret", k8SApiGroup.CreateOrUpdateSecretHandler)
}
func (*K8SRouter) initConfigMapRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/configMap/:namespace", k8SApiGroup.GetConfigMapDetailOrListHandler)
	group.DELETE("/configMap/:namespace/:name", k8SApiGroup.DeleteConfigMapHandler)
	group.POST("/configMap", k8SApiGroup.CreateOrUpdateConfigMapHandler)
}
func (*K8SRouter) initNodeRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.GET("/node", k8SApiGroup.GetNodeDetailOrListHandler)
	group.PUT("/node/label", k8SApiGroup.UpdateNodeLabelHandler)
	group.PUT("/node/taint", k8SApiGroup.UpdateNodeTaintHandler)
}
func (*K8SRouter) initPodRouter(group *gin.RouterGroup, k8SApiGroup k8s.ApiGroup) {
	group.POST("/pod", k8SApiGroup.CreateOrUpdatePodHandler)
	group.GET("/list/:namespace", k8SApiGroup.GetPodListOrDetailHandler)
	group.DELETE("/pod/:namespace/:name", k8SApiGroup.DeletePodHandler)
	group.GET("/namespace", k8SApiGroup.GetNamespaceListHandler)
}
