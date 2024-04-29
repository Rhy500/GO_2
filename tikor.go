package main

import (
	"fmt"
	"math"
)

type Titik3D struct {
	X float64
	Y float64
	Z float64
}

func main() {
	// Deklarasi variabel untuk menyimpan koordinat titik A
	var titikA Titik3D
	// Membaca input koordinat titik A
	fmt.Scan(&titikA.X, &titikA.Y, &titikA.Z)
	dimensih := Dimensi(titikA)
	// menampilkan  hasil jarak
	fmt.Println(dimensih)
}
func Dimensi(t Titik3D) float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}
