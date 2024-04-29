package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(jumlahGanjil(N))
}
func jumlahGanjil(N int) int {
	// mengembalikan jumlah 1 + 3 + 5 + ... N
	jumlah := 0
	for i := 1; i <= N; i += 2 {
		if i%2 != 0 {
			jumlah += i
		}
	}
	return jumlah
}
