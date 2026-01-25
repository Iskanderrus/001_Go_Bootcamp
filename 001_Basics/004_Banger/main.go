package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	fmt.Println(strings.Repeat("!", l) + strings.ToUpper(msg) + strings.Repeat("!", l))
}
