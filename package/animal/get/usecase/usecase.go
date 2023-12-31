package usecase

import (
	"pbf/package/animal/get/repository"
)

type UseCase struct {
	query repository.Query
}

func (a UseCase) Invoke() (Output, error) {

	entity, err := a.query.GetAnimal()
	if err != nil {
		return Output{}, err
	}

	return Output{
		Animal: entity,
	}, nil
}
