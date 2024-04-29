package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	cetak(n, m)
}
func cetak(n, m int) {
	/*
		I.S. terdefinisi dua buah bilangan bulat positif n dan m, dengan n < m
		F.S. menampilkan barisan bilangan dari n hingga m
	*/
	for i := n; i <= m; i++ {
		if i == m {
			fmt.Print(i)
		} else {
			fmt.Print(i, " ")
		}
	}
}
