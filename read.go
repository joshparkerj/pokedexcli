package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func read() {
	// the "R" in "REPL"
	fmt.Print("pokedex > ")
	scanner.Scan()
	print(scanner.Text())
}
