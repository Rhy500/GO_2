package main

import "fmt"

func urutkan(a, b *int) {

	if *a < *b {
		fmt.Println("hasil keluaran", *a, *b)
	} else {
		fmt.Println("hasil keluaran", *b, *a)
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	for x != y {
		urutkan(&x, &y)
		fmt.Scan(&x, &y)

	}

}
