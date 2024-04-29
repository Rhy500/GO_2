package main

import "fmt"

func procB(a, a int) (b, c *int) {
	b = b + c
	c = a + b + c
	return b, c
}

func main() {
	var a int
	var b, c int
	fmt.Scan(&b, &c)
	a = 10
	procB(a, a)
	fmt.Println(a)
}
