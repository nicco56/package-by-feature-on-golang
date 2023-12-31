package adaptor

import (
	"pbf/package/animal/get/usecase"

	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {

	useCaseOutput, err := usecase.UseCase{}.Invoke()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := output{
		animal: useCaseOutput.Animal,
	}

	output.generateResponse(c)
}
