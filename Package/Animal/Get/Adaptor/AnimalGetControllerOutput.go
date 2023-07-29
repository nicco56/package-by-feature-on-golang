package Adaptor

import (
	"github.com/gin-gonic/gin"

	"packageByFeature/Package/Animal/Get/Domain/Entity"
)

type AnimalGetControllerOutput struct {
	AnimalGetNameEntity Entity.AnimalGetNameEntity
}

func (out AnimalGetControllerOutput) GenerateResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": out.AnimalGetNameEntity.GetName(),
	})
}
