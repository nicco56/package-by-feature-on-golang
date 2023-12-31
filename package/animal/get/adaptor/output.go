package adaptor

import (
	"github.com/gin-gonic/gin"
	"packageByFeature/package/animal/get/entity"
)

type output struct {
	animal *entity.Animal
}

func (out output) generateResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": out.animal.GetName(),
		"age":  out.animal.GetAge(),
	})
}
