package main

import (
	"fmt"
)

// isValid checks if a given set is valid or not
func isValid(set []int) bool {
	freq := make(map[int]int)
	for _, num := range set {
		freq[num]++
		if freq[num] > 1 {
			return false
		}
	}
	return true
}

func main() {
	var a1, a2, a3, a4, a5 int
	var b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12 int
	setA := []int{a1, a2, a3, a4, a5}
	setB := []int{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12}
	fmt.Println(isValid(setA)) // true
	fmt.Println(isValid(setB)) // false
	fmt.Scan(&a1, &a2, &a3, &a4, &a5)
	fmt.Scan(&b1, &b2, &b3, &b4, &b5, &b6, &b7, &b8, &b9, &b10, &b11, &b12)
}
