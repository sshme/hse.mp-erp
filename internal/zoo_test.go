package internal

import (
	"mp-erp/internal/domain"
	"mp-erp/internal/utils"
	"testing"
)

func TestZoo_AddAnimal(t *testing.T) {
	zoo := NewZoo(NewVetClinic(utils.NewLogger()))

	rabbit, _ := domain.NewAnimal("rabbit", "Bunny", 2, 1, 7)
	zoo.AddAnimal(rabbit)

	if count := zoo.AnimalCount(); count != 1 {
		t.Errorf("AnimalCount() = %v, want %v", count, 1)
	}
}

func TestZoo_TotalFoodConsumption(t *testing.T) {
	zoo := NewZoo(NewVetClinic(utils.NewLogger()))

	rabbit, _ := domain.NewAnimal("rabbit", "Bunny", 2, 1, 7)
	tiger, _ := domain.NewAnimal("tiger", "Tiger", 10, 2, 0)

	zoo.AddAnimal(rabbit)
	zoo.AddAnimal(tiger)

	if total := zoo.TotalFoodConsumption(); total != 12 {
		t.Errorf("TotalFoodConsumption() = %v, want %v", total, 12)
	}
}

func TestZoo_ContactAnimals(t *testing.T) {
	zoo := NewZoo(NewVetClinic(utils.NewLogger()))

	kindRabbit, _ := domain.NewAnimal("rabbit", "Kind Bunny", 2, 1, 7)
	unkindRabbit, _ := domain.NewAnimal("rabbit", "Angry Bunny", 2, 2, 3)
	tiger, _ := domain.NewAnimal("tiger", "Tiger", 10, 3, 0)

	zoo.AddAnimal(kindRabbit)
	zoo.AddAnimal(unkindRabbit)
	zoo.AddAnimal(tiger)

	contactAnimals := zoo.ContactAnimals()
	if len(contactAnimals) != 1 {
		t.Errorf("ContactAnimals() returned %v animals, want 1", len(contactAnimals))
	}
}

func TestZoo_AddThing(t *testing.T) {
	zoo := NewZoo(NewVetClinic(utils.NewLogger()))

	table := &domain.Table{Thing: domain.NewThing("Table 1", 1)}
	computer := &domain.Computer{Thing: domain.NewThing("Computer 1", 2)}

	zoo.AddThing(table)
	zoo.AddThing(computer)

	items := zoo.InventoryItems()
	if len(items) != 2 {
		t.Errorf("InventoryItems() returned %v items, want 2", len(items))
	}
}

func TestZoo_ContactAnimals_WithMonkey(t *testing.T) {
	zoo := NewZoo(NewVetClinic(utils.NewLogger()))

	kindMonkey, _ := domain.NewAnimal("monkey", "Kind Monkey", 3, 1, 8)
	unkindMonkey, _ := domain.NewAnimal("monkey", "Angry Monkey", 3, 2, 4)

	zoo.AddAnimal(kindMonkey)
	zoo.AddAnimal(unkindMonkey)

	contactAnimals := zoo.ContactAnimals()
	if len(contactAnimals) != 1 {
		t.Errorf("ContactAnimals() returned %v animals, want 1", len(contactAnimals))
	}

	// Проверяем, что это именно добрая обезьяна
	if contactAnimals[0].GetName() != "Kind Monkey" {
		t.Errorf("Wrong animal in contact list: got %s, want Kind Monkey", contactAnimals[0].GetName())
	}
}
