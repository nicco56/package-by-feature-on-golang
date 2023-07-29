package UseCase

import (
	"packageByFeature/Package/Animal/Get/Domain/DB"
)

// AnimalGetUseCase AnimalGetUseCaseのストラクト(構造体)。Classみたいなもの
type AnimalGetUseCase struct {
	query DB.AnimalGetQuery
}

// Invoke 実行処理
func (a AnimalGetUseCase) Invoke() (AnimalGetUseCaseOutput, error) {

	// AnimalGetUseCase structが持っているqueryのGetAnimalメソッドを呼び出す
	entity, err := a.query.GetAnimal()

	// エラーハンドリング
	if err != nil {
		return AnimalGetUseCaseOutput{}, err
	}

	// AnimalGetUseCaseOutput に entityを渡して返す
	return AnimalGetUseCaseOutput{
		AnimalGetNameEntity: entity,
	}, nil
}

// NewAnimalGetUseCase AnimalGetUseCaseのfactoryメソッド（constructorみたいなもの）
// DI,テストで使える。
//
// 例:
// mock = Mock.SomeQuery{}
// UseCase.NewAnimalGetUseCase(mock).Invoke() みたいに使う
func NewAnimalGetUseCase(query DB.AnimalGetQuery) AnimalGetUseCase {
	return AnimalGetUseCase{
		query: query,
	}
}
