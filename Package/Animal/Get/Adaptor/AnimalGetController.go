package Adaptor

import (
	"packageByFeature/Package/Animal/Get/UseCase"

	"github.com/gin-gonic/gin"
)

func GetName(c *gin.Context) {

	// インポートしたUseCaseを呼び出す
	useCaseOutput, err := UseCase.AnimalGetUseCase{}.Invoke()

	// エラーハンドリング
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// コントローラーのOutPutを生成
	output := AnimalGetControllerOutput{
		AnimalGetNameEntity: useCaseOutput.AnimalGetNameEntity,
	}

	// レスポンス
	output.GenerateResponse(c)
}
