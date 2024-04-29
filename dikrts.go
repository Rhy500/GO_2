package main

import "fmt"

func main() {
	var X, total, digit int

	fmt.Scanln(&X)
	total = 0

	for X > 0 {
		digit = X % 10
		total = total + digit
		fmt.Print(digit, " ")
		X = X / 10
	}
	fmt.Println(" ")
	fmt.Println(total)
}
