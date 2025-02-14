package commands

import (
	"bufio"
	"mp-erp/internal"
	"mp-erp/internal/domain"
	"strings"
	"testing"
)

func TestReportInventoryCommand_Execute(t *testing.T) {
	logger := &mockLogger{}
	app := &internal.App{
		Zoo:       internal.NewZoo(internal.NewVetClinic(logger)),
		VetClinic: internal.NewVetClinic(logger),
		Logger:    logger,
	}

	rabbit, _ := domain.NewAnimal("rabbit", "Bunny", 2, 1, 7)
	app.Zoo.AddAnimal(rabbit)

	table := domain.Table{Thing: domain.Thing{}}
	app.Zoo.AddThing(&table)

	cmd := reportInventoryCommand{}
	cmd.Execute(app, bufio.NewScanner(strings.NewReader("")))

	headerFound := false
	for _, msg := range logger.messages {
		if msg == "Инвентарные предметы:" {
			headerFound = true
			break
		}
	}

	if !headerFound {
		t.Error("Inventory report header not found in output")
	}
}
