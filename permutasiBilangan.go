package main

import "fmt"

func faktorial(a int) int {
	var fakt, i int
	fakt = 1
	for i = a; i >= 1; i-- {
		fakt = fakt * i
	}
	return fakt
}

func permutasi(m, n int) float64 {
	var hasil float64
	hasil = float64(faktorial(m)) / float64(faktorial(m-n))
	return hasil
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x >= y {
		fmt.Printf("%v %v %v\n", faktorial(x), faktorial(y), permutasi(x, y))
	} else {
		fmt.Printf("%v %v %v\n", faktorial(x), faktorial(y), permutasi(y, x))
	}
}
