package animal

import (
	"github.com/gin-gonic/gin"
	AnimalGetAdaptor "pbf/package/animal/get/adaptor"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/animal", AnimalGetAdaptor.Controller)
}
