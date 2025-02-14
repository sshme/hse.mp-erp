package domain

// IPredator - интерфейс для хищных животных
type IPredator interface {
	IAlive
	IInventory
}

// Predator — implementation of IPredator
type Predator struct {
	Animal
}

func (p Predator) GetFood() int {
	return p.foodKg
}

func (p Predator) GetName() string {
	return p.name
}

func (p Predator) GetNumber() int {
	return p.number
}
