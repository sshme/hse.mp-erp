package domain

type IHerbo interface {
	IAnimal
	CanContact() bool
}

// Herbo – herbivores animals.
type Herbo struct {
	Animal
	kindness int
}

func (h Herbo) CanContact() bool {
	return h.kindness > 5
}
