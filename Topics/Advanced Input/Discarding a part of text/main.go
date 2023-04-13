package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const discardedBytes = 10

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, err := reader.Discard(discardedBytes)

	// DO NOT delete or modify the code below:
	if err != nil {
		fmt.Println(err)
	}

	data, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	fmt.Println(data)
}
