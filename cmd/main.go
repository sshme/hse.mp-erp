package main

import (
	"bufio"
	"fmt"
	"log"
	"mp-erp/internal/commands"
	"mp-erp/internal/di"
	"os"
)

func main() {

	app, err := di.InitializeApp()
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}

	app.Logger.Println("App started")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Type help to see available commands or exit to quit")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "exit" {
			break
		}

		command := commands.Get(choice)
		if command == nil {
			app.Logger.Printf("Unknown command: %s\n", choice)
			continue
		}

		command.Execute(app, scanner)
	}

	app.Logger.Println("Finished.")
}
