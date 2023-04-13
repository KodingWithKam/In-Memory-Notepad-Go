package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Which function from the `strings` package is used to replace all appearances
// of a word or character in a string with another word or character?
func symbolsCounter(text string) int {
	text = strings.ReplaceAll(text, " ", "")
	return len(text)
}

// DO NOT delete or modify the contents of the main function!
func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	allLength := len(scan.Text())
	symbols := symbolsCounter(scan.Text())
	fmt.Printf("Symbols in text: %d\nSpaces in text: %d", symbols, allLength-symbols)
}
