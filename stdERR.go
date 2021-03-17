//Simple go program to play with errors
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give one or more arguments"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "This is something like daat \n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
