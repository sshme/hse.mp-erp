package commands

import (
	"bufio"
	"fmt"
	"mp-erp/internal"
	"strings"
	"testing"
)

type mockLogger struct {
	messages []string
}

func (m *mockLogger) Println(v ...interface{}) {
	m.messages = append(m.messages, fmt.Sprint(v...))
}

func (m *mockLogger) Printf(format string, v ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, v...))
}

func TestAddAnimalCommand_Execute(t *testing.T) {
	input := strings.NewReader("rabbit\nBunny\n2\n1\n7\n")
	scanner := bufio.NewScanner(input)
	logger := &mockLogger{}

	app := &internal.App{
		Zoo:       internal.NewZoo(internal.NewVetClinic(logger)),
		VetClinic: internal.NewVetClinic(logger),
		Logger:    logger,
	}

	cmd := addAnimalCommand{}
	cmd.Execute(app, scanner)

	if app.Zoo.AnimalCount() != 1 {
		t.Error("Animal was not added to zoo")
	}
}
