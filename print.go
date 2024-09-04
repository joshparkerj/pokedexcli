package main

import (
	"fmt"
)

func print(input string) {
	// The "P" in "REPL"
	output, err := eval(input)
	if err != nil {
		fmt.Println("bad input! try again pls trainer...")
	} else {
		fmt.Println(output)
	}
}
