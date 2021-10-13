package main

import (
	"os"
	"part1/translator"
)

/*
Method hit on program execution which relays command line args to translator.
*/
func main() {
	executionArgs := os.Args[1]
	translator.Start(executionArgs)
}
