package main

import (
	"os"
)

func commandExit(_ string) (result string, err error) {
	os.Exit(0)
	return
}
