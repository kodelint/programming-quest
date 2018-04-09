package main

import (
	"fmt"
	"os"
	"strconv"
)

func fbzrecur(number int) {
	if number < 3 {
		os.Exit(0)
	}
	if number%15 == 0 {
		fmt.Println("Number:", number, "is fizzbuzz")
	} else if number%5 == 0 {
		fmt.Println("Number:", number, "is buzz")
	} else if number%3 == 0 {
		fmt.Println("Number:", number, "is fizz")
	}
	fbzrecur(number - 1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		os.Exit(1)
	}
	num, _ := strconv.Atoi(os.Args[1])
	fbzrecur(num)
}
