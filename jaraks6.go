package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x float64
	y float64
}

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	a = P1.x
	b = P1.y
	c = P2.x
	d = P2.y
	P1 := Titik{a, b}
	P2 := Titik{c, d}
	jarak := jarak(P1, P2)
	fmt.Printf("Jarak antara titik P1 dan P2 adalah: %.2f\n", jarak)
}
func jarak(P1, P2 Titik) float64 {
	return Akar(math.Pow(P1.x-P2.y, 2) + (math.Pow(P1.x-P2.y, 2)))
}
func Akar(x float64) float64 {
	return math.Sqrt(x)
}
