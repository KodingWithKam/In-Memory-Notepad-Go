package main

import "fmt"

func main() {
	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
}
