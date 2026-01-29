package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "john"
	lastname := "doe"
	fmt.Printf("My name is %s %s\n", strings.ToTitle(name), strings.ToUpper(lastname))
}
