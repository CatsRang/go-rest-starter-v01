package main

import (
	"fmt"
	"os"

	"go_ex01/pkg/root"
)

func main() {
	rootCmd := root.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
