package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	// Menampilkan 11, 12, 123, ..., 1234
	for i := 1; i <= N; i++ {
		str := ""
		for j := 1; j <= i; j++ {
			str += strconv.Itoa(j)
		}
		fmt.Println(str)
	}

	// Menampilkan 1, 1, 1, ..., 1
	for i := 1; i <= N; i++ {
		fmt.Println(1)
	}
}
