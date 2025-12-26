package main

import (
	"fmt"
	"os"

	"github.com/danielgatis/go-findfont/findfont"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-findfont <font-name>")
		fmt.Println("\nExample: go-findfont Arial")
		fmt.Println("\nTo list all fonts: go-findfont --list")
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "--list" || arg == "-l" {
		fonts := findfont.List()
		for _, f := range fonts {
			fmt.Println(f)
		}
		return
	}

	fontPath, err := findfont.Find(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(fontPath)
}
