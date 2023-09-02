package example

import (
	"github.com/gin-gonic/gin"
	"k8sdashboar.com/api"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	group := r.Group("/example")
	exampleApiGroup := api.ApiGroupApp.ExampleApiGroup
	group.GET("/ping", exampleApiGroup.ExamplePing)
}
