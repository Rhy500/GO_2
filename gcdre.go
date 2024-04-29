package main

import "fmt"

/*
func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gcd(A, B))
}
func gcd(a,b int) int {
	var c int
	for b> 0 {
		c = a % b
		a = b
		b = c
	}
	return a
}
*/
func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gcd(A, B))
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
