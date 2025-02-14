package internal

import (
	"math/rand"
	"mp-erp/internal/domain"
	"mp-erp/internal/utils"
	"time"
)

// VetClinic – check health of animals
type VetClinic struct {
	Logger utils.Logger
}

func NewVetClinic(logger utils.Logger) *VetClinic {
	return &VetClinic{Logger: logger}
}

// CheckHealth – make a decision to include an animal in the zoo.
func (vc *VetClinic) CheckHealth(a domain.IInventory) bool {
	rand.Seed(time.Now().UnixNano())
	health := rand.Intn(100)
	vc.Logger.Printf("Проверка здоровья для %s: %d", a.GetName(), health)
	return health < 80
}
