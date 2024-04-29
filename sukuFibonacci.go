package main

import "fmt"

func sumFibonacci(N int) int {
	var f1, f2, f3, i, sum int
	f1 = 0
	f2 = 1
	sum = 0
	for i = 2; i <= N; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
		sum += f1
	}
	return sum
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(sumFibonacci(N))
}
