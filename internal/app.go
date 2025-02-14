package internal

import (
	"mp-erp/internal/utils"
)

type App struct {
	Zoo       *Zoo
	VetClinic *VetClinic
	Logger    utils.Logger
}
