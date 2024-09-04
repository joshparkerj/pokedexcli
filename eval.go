package main

func eval(input string) (output string, err error) {
	// The "E" in "REPL"
	output, err = commands()[input].callback()
	return
}
