package animal

import (
	"github.com/gin-gonic/gin"
	AnimalGet "packageByFeature/package/animal/get/adaptor"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/animal", AnimalGet.Controller)
}
