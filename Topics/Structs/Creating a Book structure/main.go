package main

import "fmt"

const TotalPages = 310

// Book create the Book struct here with the Title, Author and Pages fields
type Book struct {
	Title  string
	Author string
	Pages  int
}

// DO NOT delete or modify the contents of the main function!
func main() {
	theHobbit := Book{
		Title:  "The Hobbit",
		Author: "J. R. R. Tolkien",
		Pages:  TotalPages,
	}

	fmt.Printf("%#v", theHobbit)
}
