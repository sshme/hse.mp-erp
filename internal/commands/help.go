package commands

import (
	"bufio"
	"fmt"
	"mp-erp/internal"
)

type helpCommand struct{}

func (c helpCommand) Name() string {
	return "help"
}

func (c helpCommand) Description() string {
	return "describe app"
}

func (c helpCommand) Execute(app *internal.App, scanner *bufio.Scanner) {
	fmt.Println("Available commands:")
	commands := GetAll()
	for _, cmd := range commands {
		fmt.Printf("%s — %s\n", cmd.Name(), cmd.Description())
	}
}

func init() {
	Register(helpCommand{})
}
