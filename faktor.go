package main

import "fmt"

func recursive(n int, divisor int) {
	if divisor <= n {
		if n%divisor == 0 {
			fmt.Println(divisor)
			recursive(n, divisor+1)
		} else {
			recursive(n, divisor+1)
		}
	}
}
func main() {
	var N int
	fmt.Scan(&N)
	recursive(N, 1)
}
