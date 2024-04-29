package main

import "fmt"

func main() {
	var ts, tl int
	fmt.Scan(&ts, &tl)
	fmt.Println(umur(ts, tl))
}
func umur(ts, tl int) int {
	return (ts - tl)
}
