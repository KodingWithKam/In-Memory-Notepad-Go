package main

import "fmt"

func main() {
	var num1, num2, num3 float32
	fmt.Scan(&num1, &num2, &num3)

	fmt.Println(Average(num1, num2, num3))
}

func Average(numbers ...float32) float32 {
	var sum float32
	for _, num := range numbers {
		sum += num
	}

	return sum / float32(len(numbers))
}
