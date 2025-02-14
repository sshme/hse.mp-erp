package commands

import (
	"bufio"
	"mp-erp/internal"
	"mp-erp/internal/domain"
	"strings"
	"testing"
)

func TestReportAnimalCommand_Execute(t *testing.T) {
	logger := &mockLogger{}
	app := &internal.App{
		Zoo:       internal.NewZoo(internal.NewVetClinic(logger)),
		VetClinic: internal.NewVetClinic(logger),
		Logger:    logger,
	}

	rabbit, _ := domain.NewAnimal("rabbit", "Bunny", 2, 1, 7)
	tiger, _ := domain.NewAnimal("tiger", "Tiger", 10, 2, 0)
	app.Zoo.AddAnimal(rabbit)
	app.Zoo.AddAnimal(tiger)

	cmd := reportAnimalCommand{}
	cmd.Execute(app, bufio.NewScanner(strings.NewReader("")))

	expectedMessage := "В зоопарке 2 животных. Общий расход еды: 12 кг в сутки\n"
	found := false
	for _, msg := range logger.messages {
		if msg == expectedMessage {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected message not found in logger output. Want: %s", expectedMessage)
	}
}
