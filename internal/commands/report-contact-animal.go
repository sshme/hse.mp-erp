package commands

import (
	"bufio"
	"mp-erp/internal"
)

type reportContactAnimalCommand struct{}

func (c reportContactAnimalCommand) Name() string {
	return "report-contact-animal"
}

func (c reportContactAnimalCommand) Description() string {
	return "output the list of animals for the contact zoo (herbivores, kind > 5)"
}

func (c reportContactAnimalCommand) Execute(app *internal.App, scanner *bufio.Scanner) {
	contactAnimals := app.Zoo.ContactAnimals()
	app.Logger.Println("Животные, подходящие для контактного зоопарка:")
	for _, inv := range contactAnimals {
		app.Logger.Printf("№%d: %s\n", inv.GetNumber(), inv.GetName())
	}
}

func init() {
	Register(reportContactAnimalCommand{})
}
