package main

import (
	"fmt"
	"os"

	"github.com/danielgatis/go-findfont/findfont"
)

func main() {
	fonts, err := findfont.Find("Emoji", findfont.FontRegular)

	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	for _, f := range fonts {
		fmt.Printf("Family: %v\nStyle : %v\nPath  : %v\n\n", f[0], f[1], f[2])
	}
}
