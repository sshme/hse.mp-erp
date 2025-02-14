//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"mp-erp/internal"
	"mp-erp/internal/utils"
)

// InitializeApp – точка входа для создания App с помощью Wire.
func InitializeApp() (*internal.App, error) {
	wire.Build(
		utils.NewLogger,       // Logger
		internal.NewVetClinic, // VetClinic depends on Logger
		internal.NewZoo,       // Zoo depends on VetClinic
		wire.Struct(new(internal.App), "Zoo", "VetClinic", "Logger"), // собираем App
	)
	return &internal.App{}, nil
}
