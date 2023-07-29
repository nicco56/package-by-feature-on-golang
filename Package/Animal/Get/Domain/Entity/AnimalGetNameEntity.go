package Entity

type AnimalGetNameEntity struct {
	Name string
}

func NewAnimalGetNameEntity(name string) AnimalGetNameEntity {
	return AnimalGetNameEntity{Name: name}
}

func (a AnimalGetNameEntity) GetName() string {
	return a.Name
}
