package main

import (
	"fmt"
	"strconv"
)

func IsMultiple(number int, radix int) bool {
	remainder := number % radix
	return remainder == 0
}

func IsMultipleOf3Or5(number int) bool {
	return IsMultiple(number, 3) || IsMultiple(number, 5)
}

func GetSumOfMultiples(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if IsMultipleOf3Or5(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	sum := GetSumOfMultiples(1000)
	fmt.Printf(strconv.Itoa(sum) + "\n")
}
