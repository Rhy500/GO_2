package main

import "fmt"

func cetakRerata(N int, i int, jumlah int, rata float64) {
	if i <= N {
		jumlah += i
		rata = float64(jumlah) / float64(i)
		fmt.Println(jumlah, rata)
		cetakRerata(N, i+1, jumlah, rata)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	cetakRerata(N, 1, 0, 0)
}
