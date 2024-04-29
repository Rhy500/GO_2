package main

import "fmt"

func main() {
	var x, y float32
	var hasil string
	fmt.Scan(&x, &y)
	fmt.Println(kuadran(x, y, hasil))
}
func kuadran(x, y float32, hasil string) string {
	if x >= 0 && y >= 0 {
		hasil = "Kuadran I"
	} else if x <= 0 && y >= 0 {
		hasil = "kuadrat II"
	} else if x <= 0 && y <= 0 {
		hasil = "kuadrat III"
	} else {
		hasil = "kuadrat IV"
	}
	return hasil
}
