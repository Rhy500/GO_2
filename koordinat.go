package main

import (
	"fmt"
	"math"
)

func main() {
	var R, T float64
	fmt.Scan(&R, &T)
	lengthX := panjangX(R, T)
	lengthY := panjangY(R, T)
	fmt.Println(lengthX)
	fmt.Println(lengthY)
}
func panjangX(R, T float64) float64 {
	return R * math.Cos(T*math.Pi/180)
}
func panjangY(R, T float64) float64 {
	return R * math.Sin(T*math.Pi/180)
}
