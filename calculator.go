package main

import (
	"fmt"
	"github.com/joaopapereira/versioned-module/v3/pkg/calculator"
)

func main() {
	fmt.Println("Going to do some math")
	calculator.AddString("1", "1")

	fmt.Println("Now for some sub")
	calculator.SubString("2", "1")
}
