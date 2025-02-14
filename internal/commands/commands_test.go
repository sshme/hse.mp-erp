package commands

import (
	"bufio"
	"mp-erp/internal"
	"testing"
)

type testCommand struct {
	name string
	desc string
}

func (c testCommand) Name() string {
	return c.name
}

func (c testCommand) Description() string {
	return c.desc
}

func (c testCommand) Execute(*internal.App, *bufio.Scanner) {

}

func TestRegisterAndGet(t *testing.T) {
	Registry = make(map[string]Command)

	cmd1 := testCommand{name: "test1", desc: "Test Command 1"}
	cmd2 := testCommand{name: "test2", desc: "Test Command 2"}

	Register(cmd1)
	if len(Registry) != 1 {
		t.Errorf("Register() failed, expected registry size 1, got %d", len(Registry))
	}

	got := Get("test1")
	if got == nil {
		t.Error("Get() returned nil for existing command")
		return
	}
	if got.Name() != "test1" || got.Description() != "Test Command 1" {
		t.Errorf("Get() returned wrong command: got %v, want %v", got, cmd1)
	}

	got = Get("nonexistent")
	if got != nil {
		t.Error("Get() should return nil for nonexistent command")
	}

	newCmd1 := testCommand{name: "test1", desc: "New Test Command 1"}
	Register(newCmd1)
	if len(Registry) != 1 {
		t.Errorf("Register() failed on overwrite, expected registry size 1, got %d", len(Registry))
	}
	got = Get("test1")
	if got.Description() != "New Test Command 1" {
		t.Errorf("Register() failed to overwrite command, got description %s, want %s",
			got.Description(), "New Test Command 1")
	}

	Register(cmd2)
}

func TestGetAll(t *testing.T) {
	Registry = make(map[string]Command)

	cmd1 := testCommand{name: "test1", desc: "Test Command 1"}
	cmd2 := testCommand{name: "test2", desc: "Test Command 2"}
	cmd3 := testCommand{name: "test3", desc: "Test Command 3"}

	Register(cmd1)
	Register(cmd2)
	Register(cmd3)

	commands := GetAll()

	if len(commands) != 3 {
		t.Errorf("GetAll() returned wrong number of commands: got %d, want 3", len(commands))
	}

	foundCommands := make(map[string]bool)
	for _, cmd := range commands {
		foundCommands[cmd.Name()] = true
	}

	expectedCommands := []string{"test1", "test2", "test3"}
	for _, name := range expectedCommands {
		if !foundCommands[name] {
			t.Errorf("GetAll() missing command: %s", name)
		}
	}
}

func TestRegisterDuplicateNames(t *testing.T) {
	Registry = make(map[string]Command)

	cmd1 := testCommand{name: "test", desc: "Original Command"}
	cmd2 := testCommand{name: "test", desc: "Duplicate Command"}

	Register(cmd1)
	Register(cmd2)

	if len(Registry) != 1 {
		t.Errorf("Registry should contain only one command, got %d", len(Registry))
	}

	got := Get("test")
	if got.Description() != "Duplicate Command" {
		t.Errorf("Expected duplicate command to override original, got description %s", got.Description())
	}
}

func TestEmptyRegistry(t *testing.T) {
	Registry = make(map[string]Command)

	if got := Get("anything"); got != nil {
		t.Error("Get() should return nil for empty registry")
	}

	if got := GetAll(); len(got) != 0 {
		t.Error("GetAll() should return empty slice for empty registry")
	}
}
