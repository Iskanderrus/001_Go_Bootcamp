package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Give me args")
	} else if len(args) == 1 {
		fmt.Println("There is one:", args[0])
	} else if len(args) == 2 {
		fmt.Println("There are two:", args[0], args[1])
	} else {
		fmt.Printf("There are %v arguments\n", len(args))
	}
}
