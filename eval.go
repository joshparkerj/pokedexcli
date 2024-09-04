package main

import (
	"errors"
)

func eval(input string) (output string, err error) {
	// The "E" in "REPL"
	command, ok := commands()[input]
	if !ok {
		err = errors.New("command not found")
	} else {
		output, err = command.callback()
	}

	return
}
