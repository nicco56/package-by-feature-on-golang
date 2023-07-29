package UseCase

import (
	"packageByFeature/Package/Animal/Get/Domain/DB"
)

type AnimalGetUseCase struct {
	query DB.AnimalGetQuery
}

func (a AnimalGetUseCase) Invoke() (AnimalGetUseCaseOutput, error) {
	entity, err := a.query.GetAnimal()
	if err != nil {
		return AnimalGetUseCaseOutput{}, err
	}

	return AnimalGetUseCaseOutput{
		AnimalGetNameEntity: entity,
	}, nil
}

func NewAnimalGetUseCase(query DB.AnimalGetQuery) AnimalGetUseCase {
	return AnimalGetUseCase{
		query: query,
	}
}
