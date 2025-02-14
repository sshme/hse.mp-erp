package internal

import "mp-erp/internal/domain"

// Zoo – aggregator for managing animals and things. It has a dependency on the veterinary clinic.
type Zoo struct {
	animals    []domain.IInventory // animals implement IInventory (so we can uniformly get their name and number)
	things     []domain.IInventory
	veterinary *VetClinic
}

func NewZoo(vet *VetClinic) *Zoo {
	return &Zoo{
		animals:    make([]domain.IInventory, 0),
		things:     make([]domain.IInventory, 0),
		veterinary: vet,
	}
}

func (z *Zoo) AddAnimal(a domain.IInventory) {
	z.animals = append(z.animals, a)
}

func (z *Zoo) AddThing(t domain.IInventory) {
	z.things = append(z.things, t)
}

func (z *Zoo) AnimalCount() int {
	return len(z.animals)
}

func (z *Zoo) TotalFoodConsumption() int {
	total := 0
	for _, a := range z.animals {
		if alive, ok := a.(domain.IAlive); ok {
			total += alive.GetFood()
		}
	}
	return total
}

// ContactAnimals – selects animals that can be interacted with.
func (z *Zoo) ContactAnimals() []domain.IInventory {
	contactList := make([]domain.IInventory, 0)
	for _, a := range z.animals {
		switch obj := a.(type) {
		case domain.Rabbit:
			if obj.CanContact() {
				contactList = append(contactList, a)
			}
		case domain.Monkey:
			if obj.CanContact() {
				contactList = append(contactList, a)
			}
		}
	}
	return contactList
}

// InventoryItems – returns all items in the zoo that implement the IInventory interface.
func (z *Zoo) InventoryItems() []domain.IInventory {
	all := make([]domain.IInventory, 0)
	all = append(all, z.animals...)
	all = append(all, z.things...)
	return all
}
