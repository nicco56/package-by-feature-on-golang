package Adaptor

import (
	"packageByFeature/Package/Animal/Get/UseCase"

	"github.com/gin-gonic/gin"
)

func GetName(c *gin.Context) {
	useCaseOutput, err := UseCase.AnimalGetUseCase{}.Invoke()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := AnimalGetControllerOutput{
		AnimalGetNameEntity: useCaseOutput.AnimalGetNameEntity,
	}
	output.GenerateResponse(c)
}
