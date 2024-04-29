package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= 10; i++ {
		fmt.Printf("S%d = %d\n", i, fibonacci(i))
	}
}
