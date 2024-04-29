package main

import (
	"fmt"
	"math"
)

func calculateDistance(x1, y1, z1, x2, y2, z2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2) + math.Pow(z2-z1, 2))
}

func main() {
	var xFalcon, yFalcon, zFalcon float64
	var xA, yA, zA float64
	var xB, yB, zB float64
	var planet string

	// Input posisi pesawat Millenium Falcon, posisi planet A, posisi planet B, dan pilihan planet
	fmt.Scan(&xFalcon, &yFalcon, &zFalcon)
	fmt.Scan(&xA, &yA, &zA)
	fmt.Scan(&xB, &yB, &zB)
	fmt.Scan(&planet)

	// Hitung jarak pesawat ke masing-masing planet
	distanceToA := calculateDistance(xFalcon, yFalcon, zFalcon, xA, yA, zA)
	distanceToB := calculateDistance(xFalcon, yFalcon, zFalcon, xB, yB, zB)

	// Tentukan planet terdekat
	var nearestPlanet string
	if distanceToA < distanceToB {
		nearestPlanet = "A"
	} else {
		nearestPlanet = "B"
	}

	// Output hasil
	fmt.Printf("%.2f %s\n", math.Abs(distanceToA), planet)
}
