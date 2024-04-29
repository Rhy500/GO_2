package main

import "fmt"

func main() {
	var n, luas, keliling, panjang, lebar int
	fmt.Scan(&n)
	keliling = 0
	luas = 0
	for i := n; i >= n; i++ {
		fmt.Scan(&panjang)
		fmt.Scan(&lebar)
		luas = panjang * lebar
		keliling = 2 * (panjang + lebar)
		fmt.Println(luas, keliling)
	}

}
