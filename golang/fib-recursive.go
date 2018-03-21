package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
  number := 10
  fmt.Println("Fibonacci Sequence from 0 to", number)
  fmt.Println("================================")
	for i := 0; i <= number; i++ {
		fmt.Println(fib(i))
	}
}
