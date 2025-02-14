package commands

import (
	"bufio"
	"mp-erp/internal"
	"mp-erp/internal/domain"
	"strings"
	"testing"
)

func TestReportContactAnimalCommand_Execute(t *testing.T) {
	logger := &mockLogger{}
	app := &internal.App{
		Zoo:       internal.NewZoo(internal.NewVetClinic(logger)),
		VetClinic: internal.NewVetClinic(logger),
		Logger:    logger,
	}

	kindRabbit, _ := domain.NewAnimal("rabbit", "Kind Bunny", 2, 1, 7)
	unkindRabbit, _ := domain.NewAnimal("rabbit", "Angry Bunny", 2, 2, 3)
	kindMonkey, _ := domain.NewAnimal("monkey", "Kind Monkey", 3, 3, 8)
	tiger, _ := domain.NewAnimal("tiger", "Tiger", 10, 4, 0)

	app.Zoo.AddAnimal(kindRabbit)
	app.Zoo.AddAnimal(unkindRabbit)
	app.Zoo.AddAnimal(kindMonkey)
	app.Zoo.AddAnimal(tiger)

	cmd := reportContactAnimalCommand{}
	cmd.Execute(app, bufio.NewScanner(strings.NewReader("")))

	headerFound := false
	kindBunnyFound := false
	kindMonkeyFound := false

	for _, msg := range logger.messages {
		if msg == "Животные, подходящие для контактного зоопарка:" {
			headerFound = true
		}
		if strings.Contains(msg, "Kind Bunny") {
			kindBunnyFound = true
		}
		if strings.Contains(msg, "Kind Monkey") {
			kindMonkeyFound = true
		}
	}

	if !headerFound {
		t.Error("Contact animals report header not found")
	}
	if !kindBunnyFound {
		t.Error("Kind Bunny not found in contact animals")
	}
	if !kindMonkeyFound {
		t.Error("Kind Monkey not found in contact animals")
	}
}
