package domain

import (
	"errors"
	"fmt"
	"strings"
)

// IAlive – interface for living things
type IAlive interface {
	GetFood() int
}

// IInventory – interface for inventory items
type IInventory interface {
	GetNumber() int
	GetName() string
}

type IAnimal interface {
	IAlive
	IInventory
}

// Animal – базовый тип для животных; реализует IAlive и IInventory.
type Animal struct {
	name   string
	foodKg int
	number int
}

func (a Animal) GetFood() int {
	return a.foodKg
}

func (a Animal) GetNumber() int {
	return a.number
}

func (a Animal) GetName() string {
	return a.name
}

// NewAnimal – fabric method for creating new animals.
// For herbivores (rabbit, monkey) the kindness parameter is expected, for predators it is ignored.
func NewAnimal(animalType, name string, foodKg, number, kindness int) (IInventory, error) {
	switch strings.ToLower(animalType) {
	case "rabbit":
		return Rabbit{Herbo{Animal{name, foodKg, number}, kindness}}, nil
	case "monkey":
		return Monkey{Herbo{Animal{name, foodKg, number}, kindness}}, nil
	case "tiger":
		return Tiger{Predator{Animal{name, foodKg, number}}}, nil
	case "wolf":
		return Wolf{Predator{Animal{name, foodKg, number}}}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported animal type: %s", animalType))
	}
}
