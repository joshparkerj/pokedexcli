package main

import (
	"errors"
	"strings"
)

func eval(input string) (output string, err error) {
	// The "E" in "REPL"
	inputs := strings.SplitN(input, " ", 2)
	command, ok := commands()[inputs[0]]
	if !ok {
		err = errors.New("command not found")
	} else if len(inputs) == 2 {
		output, err = command.callback(inputs[1])
	} else {
		output, err = command.callback("")
	}

	return
}
