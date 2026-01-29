package main

import "fmt"

func main() {
	score, valid := 2, true

	if score > 3 && valid {
		fmt.Println("Good!")
	} else if score < 3 && valid {
		fmt.Println("Average!")
	} else {
		fmt.Println("Not good enough!")
	}
}
