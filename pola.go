package main

import "fmt"

func main() {
	var alas, i, j int

	fmt.Scan(&alas)

	for i = 0; i < alas; i++ {
		for j = 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
