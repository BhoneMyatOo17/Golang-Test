package main

import (
	"errors"
	"fmt"
)

func main() {
	var printText string = "Hey Girl Hey"
	printme(printText)

	var numerator int = 12
	var denominator int = 0
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Printf(" %v & %v", result, remainder)
}

func printme(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {

	if denominator == 0 {
		err := errors.New("Cannot divide by 0")
		fmt.Println(err)
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
