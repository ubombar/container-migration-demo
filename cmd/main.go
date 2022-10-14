package main

import (
	"fmt"
	"strconv"

	"github.com/ubombar/container-migration-demo/pkg/app"
)

func displayMenu() {
	fmt.Println("'m' and '<int>' (migrate the process with specified pid) 'q' (quit) :")
}

func migrateProcess(pif uint32) {

}

func main() {
	appClient := app.NewClient()
	var exitRequested bool = false

	for !exitRequested {
		displayMenu()

		input := appClient.GetInput()

		switch input {
		case "q":
			exitRequested = true
		case "m":
			pidText := appClient.GetInput()

			pid, err := strconv.Atoi(pidText)

			if err == nil {
				appClient.MigrateContainer(int32(pid))
			} else {
				fmt.Println("Cannot cast the input to integer try again")
			}
		}
	}

}
