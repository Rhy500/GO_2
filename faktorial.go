package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	faktorialX := hitungFaktorial(x)
	faktorialY := hitungFaktorial(y)
	permutasi := hitungPermutasi(x, y)
	fmt.Printf("%d\n", faktorialX, faktorialY, permutasi)
}
func hitungFaktorial(x int) int {
	result := 1
	for i := 2; i <= x; i++ {
		result *= i
	}
	return result
}
func hitungPermutasi(x, y int) int {
	if x > y {
		return hitungFaktorial(x) / hitungFaktorial(x-y)
	} else {
		return hitungFaktorial(y) / hitungFaktorial(y-x)
	}
}
