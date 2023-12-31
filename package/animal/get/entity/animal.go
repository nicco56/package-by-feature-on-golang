package entity

type Animal struct {
	name string
	age  uint8
}

func NewAnimal(name string, age uint8) *Animal {
	return &Animal{name: name, age: age}
}

// GetName Go界隈ではgetterは非推奨。確かにオーバーすぎるな。
func (a Animal) GetName() string {
	return a.name
}

// GetAge Go界隈ではgetterは非推奨。確かにオーバーすぎるな。
func (a Animal) GetAge() uint8 {
	return a.age
}
