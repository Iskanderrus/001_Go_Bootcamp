package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)

	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(n)
}
