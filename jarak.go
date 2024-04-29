package main 

import (
	"fmt"
	"math"
)
type Titik struct {
	X, Y float64
}

func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	var titik1, titik2 Titik

	fmt.Print("Masukkan koordinat titik 1 (x, y): ")
	fmt.Scan(&titik1.X, &titik1.Y)

	fmt.Print("Masukkan koordinat titik 2 (x, y): ")
	fmt.Scan(&titik2.X, &titik2.Y)

	hasil := jarak(titik1, titik2)
	fmt.Printf("Jarak antara titik 1 dan titik 2 adalah: %.2f\n", hasil)
}