package main

import (
	"os"
)

func commandExit() (result string, err error) {
	os.Exit(0)
	return
}
