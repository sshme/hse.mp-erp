package commands

import (
	"bufio"
	"mp-erp/internal"
)

type Command interface {
	Name() string
	Description() string
	Execute(*internal.App, *bufio.Scanner)
}

var Registry = make(map[string]Command)

func Get(name string) Command {
	return Registry[name]
}

func GetAll() []Command {
	commands := make([]Command, 0, len(Registry))
	for _, cmd := range Registry {
		commands = append(commands, cmd)
	}
	return commands
}

func Register(cmd Command) {
	Registry[cmd.Name()] = cmd
}
