package DB

import "packageByFeature/Package/Animal/Get/Domain/Entity"

type AnimalGetQuery struct{}

func (q AnimalGetQuery) GetAnimal() (Entity.AnimalGetNameEntity, error) {
	// ここでDBからデータを取得する処理
	name, err := "DOG", error(nil)

	// エラーハンドリング
	if err != nil {
		return Entity.AnimalGetNameEntity{}, err
	}

	return Entity.NewAnimalGetNameEntity(name), nil
}
