package main

import "fmt"

func bintang(n int, b string) string {
	if n <= 1 {
		return b
	} else {
		return bintang(n-1, b+"*")
	}
}

func pola(n int, a, b string) string {
	if n <= 0 {
		return ""
	} else {
		fmt.Println(a, bintang(n, ""))
		return pola(n-1, a+"*", bintang(n-1, ""))
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n, "", "")
}
