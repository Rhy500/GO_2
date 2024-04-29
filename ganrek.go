package main

import "fmt"

func bilanganGanjilRekursif(n int) {
	if n <= 1 {
		if n%2 == 1 {
			fmt.Println(n)
		}
		bilanganGanjilRekursif(n + 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(bilanganGanjilRekursif(n))
}
