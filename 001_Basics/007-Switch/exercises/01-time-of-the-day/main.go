package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	switch now.Hour() {
	case 6, 7, 8, 9, 10, 11:
		fmt.Println("Good morning!")
	case 12, 13, 14, 15, 16, 17:
		fmt.Println("Good afternoon!")
	case 18, 19, 20, 21, 22, 23:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}
}
