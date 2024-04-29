package main

import "fmt"

type Rectangle struct {
	Length int
	Width  int
	Color  string
	Geometry
}

type Geometry struct {
	Area      int
	Perimeter int
}

func isiData(persegi *Rectangle) {
	fmt.Scan(&persegi.Length, &persegi.Width, &persegi.Color)
}

func hitung(persegi *Rectangle) {
	persegi.Geometry.Area = persegi.Length * persegi.Width
	persegi.Geometry.Perimeter = 2 * (persegi.Length + persegi.Width)
}

func main() {
	var persegi Rectangle
	isiData(&persegi)
	hitung(&persegi)
	fmt.Printf("Luas: %d, Keliling: %d\n", persegi.Geometry.Area, persegi.Geometry.Perimeter)
}
