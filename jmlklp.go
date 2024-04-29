package main

import "fmt"

func main() {
	var bil1, bil2, x int
	fmt.Scan(&bil1, &bil2, &x)
	fmt.Println(jumlahKelipatan(bil1, bil2, x))
}
func jumlahKelipatan(bil1, bil2, x int) int {
	sum := 0
	for i := bil1; i <= bil2; i++ {
		if i%x == 0 {
			sum += i
		}
	}
	return sum
}
