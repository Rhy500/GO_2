package main

import "fmt"

func main() {
	var banyakper, baris, jumlah, i int
	fmt.Scanln(&banyakper)
	jumlah = 0

	for i = 1; i <= banyakper; i++ {
		fmt.Scanln(&baris)
		jumlah = jumlah + baris
	}

}
