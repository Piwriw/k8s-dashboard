package initiallize

import (
	"github.com/gin-gonic/gin"
	"k8sdashboar.com/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	exampleRouter := router.RouterGroupApp.ExampleRouterApp
	exampleRouter.InitExample(r)
	k8sRouter := router.RouterGroupApp.InitK8SRouterApp
	k8sRouter.InitK8SRouter(r)
	return r
}
