package main

import (
	"fmt"
)

func getNthNumber(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	} else if n == 10 {
		//return 2*getNthNumber(n-1) + getNthNumber(n-2)
		return getNthNumber(n*n) + getNthNumber(n/2)
	} else {
		return getNthNumber(n-1) + 2*getNthNumber(n-2)
	}
	//if n%2 == 0 {
	//return getNthNumber(n-2) + 2*getNthNumber(n-1)
	//} //else {
	//return 2*getNthNumber(n-1) + getNthNumber(n-2)
	//}
}

func main() {
	var n int
	fmt.Scan(&n)
	result := getNthNumber(n)
	fmt.Println(result)
}
