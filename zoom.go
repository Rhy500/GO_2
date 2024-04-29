package main

import "fmt"

func main() {
	var w, h, s int
	fmt.Scan(&w, &h, &s)
	if s >= 0 {
		fmt.Println(zoomIn(w, h, s))
	} else {
		fmt.Println((zoomOut(w, h, s)))
	}
}
func zoomIn(w, h, s int) (neww, newh int) {
	neww = w * s
	newh = h * s
	return neww, newh
}
func zoomOut(w, h, s int) (neww, newh int) {
	neww = (w / s) * -1
	newh = (h / s) * -1
	return neww, newh
}
