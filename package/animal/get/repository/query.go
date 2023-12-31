package repository

import (
	"packageByFeature/package/animal/get/entity"
)

type Query struct{}

func (q Query) GetAnimal() (*entity.Animal, error) {

	name := "DOG"
	age := uint8(5)

	var err error = nil

	if err != nil {
		return nil, err
	}

	return entity.NewAnimal(name, age), nil
}
