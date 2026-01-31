package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user     = "jack"
	user2    = "inanc"
)

var (
	pass  = os.Getenv("JACK_PASSWORD")
	pass2 = os.Getenv("INANC_PASSWORD")
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch u, p := args[1], args[2]; {
	case u == user && p == pass || u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
