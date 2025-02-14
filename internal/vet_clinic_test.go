package internal

import (
	"mp-erp/internal/domain"
	"mp-erp/internal/utils"
	"testing"
)

func TestVetClinic_CheckHealth(t *testing.T) {
	vet := NewVetClinic(utils.NewLogger())
	animal, _ := domain.NewAnimal("rabbit", "Test Rabbit", 2, 1, 7)

	// Since CheckHealth uses random, we can test multiple times to ensure it returns both true and false
	results := make(map[bool]bool)
	for i := 0; i < 100; i++ {
		result := vet.CheckHealth(animal)
		results[result] = true
	}

	if len(results) != 2 {
		t.Error("CheckHealth() should return both true and false over multiple calls")
	}
}
