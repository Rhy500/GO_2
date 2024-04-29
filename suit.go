package main

import "fmt"

func main() {
	var n, A, B int
	fmt.Scan(&n)
	gameSuit(n, &A, &B)
	fmt.Println(A, B)
}
func gameSuit(n int, A, B *int) {
	*A = 0
	*B = 0
	for i := 0; i < n; i++ {
		var a, b byte
		fmt.Scan(&a, &b)
		if a == 'g' && b == 'k' || a == 'k' && b == 'b' || a == 'b' && b == 'g' {
			(*A)++
		} else if a == b {
			// draw
		} else {
			(*B)++
		}
	}
}
