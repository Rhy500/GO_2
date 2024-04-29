package main

import (
	"fmt"
	"math"
)

func konversiDerajatkeRadian(T float64) float64 {
	return T * math.Pi / 180
}

func waktu(V, T float64) float64 {
	return (V * math.Sin(konversiDerajatkeRadian(T))) / 9.8
}

func jarak(V, T float64) float64 {
	return (V * V * math.Sin(2*konversiDerajatkeRadian(T))) / 9.8
}

func ketinggian(V, T float64) float64 {
	return (V * V * math.Pow(math.Sin(konversiDerajatkeRadian(T)), 2)) / (2 * 9.8)
}

func main() {
	var V, T float64
	fmt.Scan(&V, &T)
	fmt.Printf("%.2f %.2f %.2f\n", round(waktu(V, T), 2), round(jarak(V, T), 2), round(ketinggian(V, T), 2))
}

func round(num float64, ndigits int) float64 {
	pow := math.Pow(10, float64(ndigits))
	return math.Round(num*pow) / pow
}
