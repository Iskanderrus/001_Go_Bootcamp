package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	max := os.Args[1]
	n, err := strconv.Atoi(max)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}
	fmt.Printf("%5s", " ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%5d", i)
		for j := 1; j <= n; j++ {
			if j == i {
				fmt.Printf("%5d", i*j)
			} else {
				fmt.Printf("%5s", " ")
			}
		}
		fmt.Println()
	}
}
