package main

import (
	"os"

	"github.com/hjhsamuel/zerostack/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
