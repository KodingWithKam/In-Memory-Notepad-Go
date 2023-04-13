package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FloatPrecision = 64

// Do not delete the ScanFloats() function!
func ScanFloats(data []byte, atEOF bool) (int, []byte, error) {
	advance, token, err := bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseFloat(string(token), FloatPrecision)
	}
	return advance, token, err
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	bufio.s
	// Set 'ScanFloats' as the split function below
	scanner.Split(ScanFloats)

	// Write the for loop to scan and then output the scanned data
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
