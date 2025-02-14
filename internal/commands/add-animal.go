package commands

import (
	"bufio"
	"fmt"
	"mp-erp/internal"
	"mp-erp/internal/domain"
	"strconv"
	"strings"
)

type addAnimalCommand struct{}

func (c addAnimalCommand) Name() string {
	return "add"
}

func (c addAnimalCommand) Description() string {
	return "new animal"
}

func (c addAnimalCommand) Execute(app *internal.App, scanner *bufio.Scanner) {
	app.Logger.Println("Добавление нового животного")
	fmt.Print("Введите тип животного (rabbit, tiger, wolf, monkey): ")
	scanner.Scan()
	animalType := strings.ToLower(strings.TrimSpace(scanner.Text()))
	fmt.Print("Введите наименование животного: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Введите количество потребляемой еды (кг в сутки): ")
	scanner.Scan()
	foodStr := scanner.Text()
	food, err := strconv.Atoi(foodStr)
	if err != nil {
		app.Logger.Println("Неверное значение для еды")
		return
	}
	fmt.Print("Введите инвентаризационный номер: ")
	scanner.Scan()
	numberStr := scanner.Text()
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		app.Logger.Println("Неверное значение для инвентаризационного номера")
		return
	}

	var kindness int
	if animalType == "rabbit" || animalType == "monkey" {
		fmt.Print("Введите уровень доброты (0-10): ")
		scanner.Scan()
		kindStr := scanner.Text()
		kindness, err = strconv.Atoi(kindStr)
		if err != nil {
			app.Logger.Println("Неверное значение для доброты")
			return
		}
	}

	newAnimal, err := domain.NewAnimal(animalType, name, food, number, kindness)
	if err != nil {
		app.Logger.Println("Ошибка создания животного:" + err.Error())
		return
	}

	approved := app.VetClinic.CheckHealth(newAnimal)
	if !approved {
		app.Logger.Println("Животное не прошло проверку ветеринарной клиники.")
		return
	}
	app.Zoo.AddAnimal(newAnimal)
	app.Logger.Println("Животное успешно добавлено в зоопарк.")
}

func init() {
	Register(addAnimalCommand{})
}
