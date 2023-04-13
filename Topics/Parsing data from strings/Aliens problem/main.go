package main

import (
	"fmt"
	"strconv"
)

const bitSize = 64

func main() {
	var num string
	fmt.Scanln(&num)
	var base int
	fmt.Scanln(&base)

	dec, err := strconv.ParseInt(num, base, bitSize)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Num %s, base %d is %d in decimal", num, base, dec)
}
