package commands

import (
	"bufio"
	"mp-erp/internal"
)

type reportAnimalCommand struct{}

func (c reportAnimalCommand) Name() string {
	return "report-animal"
}

func (c reportAnimalCommand) Description() string {
	return "output a report on the number of animals and consumed kg of food per day"
}

func (c reportAnimalCommand) Execute(app *internal.App, scanner *bufio.Scanner) {
	count := app.Zoo.AnimalCount()
	totalFood := app.Zoo.TotalFoodConsumption()
	app.Logger.Printf("В зоопарке %d животных. Общий расход еды: %d кг в сутки\n", count, totalFood)
}

func init() {
	Register(reportAnimalCommand{})
}
