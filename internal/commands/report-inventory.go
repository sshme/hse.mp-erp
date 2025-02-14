package commands

import (
	"bufio"
	"mp-erp/internal"
)

type reportInventoryCommand struct{}

func (c reportInventoryCommand) Name() string {
	return "report-inventory"
}

func (c reportInventoryCommand) Description() string {
	return "output the list of inventory items"
}

func (c reportInventoryCommand) Execute(app *internal.App, scanner *bufio.Scanner) {
	items := app.Zoo.InventoryItems()
	app.Logger.Println("Инвентарные предметы:")
	for _, inv := range items {
		app.Logger.Printf("№%d: %s\n", inv.GetNumber(), inv.GetName())
	}
}

func init() {
	Register(reportInventoryCommand{})
}
