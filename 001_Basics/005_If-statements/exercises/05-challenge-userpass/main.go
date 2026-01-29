package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: [username] [password]")
	} else if args[0] == "jack" && args[1] == "1888" {
		fmt.Println("Access granted to \"jack\".")
	} else if args[0] == "jack" && args[1] != "1888" {
		fmt.Println("Invalid password for \"jack\".")
	} else {
		fmt.Printf("Access denied for %s.\n", args[0])
	}
}
