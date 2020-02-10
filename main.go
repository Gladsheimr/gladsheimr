package main

import (
	"os"

	"github.com/gladsheimr/gladsheimr/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
