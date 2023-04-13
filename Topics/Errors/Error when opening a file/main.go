package main

import (
	"fmt"
	"os"
)

func checkFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}

// DO NOT delete or modify the contents of the main function!
func main() {
	checkFile("nonexisting_file.txt")
}
