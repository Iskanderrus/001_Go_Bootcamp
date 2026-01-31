package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: [command] [string]")
		return
	}

	switch args[0] {
	case "lower":
		fmt.Println(strings.ToLower(args[1]))
	case "upper":
		fmt.Println(strings.ToUpper(args[1]))
	case "title":
		fmt.Println(cases.Title(language.English).String(args[1]))
	default:
		fmt.Printf("Unknown command: %q\n", args[0])
	}
}
