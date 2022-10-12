package main

import (
	"fmt"

	"github.com/checkpoint-restore/go-criu/v6"
)

func main() {
	c := criu.MakeCriu()

	version, err := c.GetCriuVersion()

	if err != nil {
		fmt.Println("criu cannot found")
		return
	}

	fmt.Printf("criu version: %v\n", version)
}
