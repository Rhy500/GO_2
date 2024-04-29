package main

import "fmt"

const N = 1000

func lebihBesar(T [N]int, jumlah, target int) int {
	total := 0
	for i := 0; i < jumlah; i++ {
		if T[i] > target {
			total += T[i]
		}
	}
	return total
}

func main() {
	var T [N]int
	var jumlah, target int
	fmt.Scanln(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&target)
	fmt.Println(lebihBesar(T, jumlah, target))
}
