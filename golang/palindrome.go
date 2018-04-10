package main

import (
	"fmt"
	"os"
)

func isPalindrom(s string) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	if s == string(runes) {
		fmt.Println("Given String:", s, "is palindrom")
	} else {
		fmt.Println("Given String:", s, "is not palindrom")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a String")
		os.Exit(1)
	}
	isPalindrom(os.Args[1])
}
