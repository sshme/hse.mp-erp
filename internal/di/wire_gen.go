// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"mp-erp/internal"
	"mp-erp/internal/utils"
)

// Injectors from wire.go:

// InitializeApp – точка входа для создания App с помощью Wire.
func InitializeApp() (*internal.App, error) {
	logger := utils.NewLogger()
	vetClinic := internal.NewVetClinic(logger)
	zoo := internal.NewZoo(vetClinic)
	app := &internal.App{
		Zoo:       zoo,
		VetClinic: vetClinic,
		Logger:    logger,
	}
	return app, nil
}
