package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "C:/Users/User/Desktop/test.txt"
	directory, fileName := path.Split(filePath)

	fmt.Println("File name: ", fileName)
	fmt.Println("Directory: ", directory)
}
