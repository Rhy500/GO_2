package main

import (
	"fmt"
	"math"
)

type data struct {
	R, S, LL, LP, KL, KP, TL, TP float64
}

type arrdata [100]data

func main() {
	var s, l float64
	var i int = 0
	var array arrdata
	fmt.Scan(&l, &s)
	for s != 0 && l != 0 {
		array[i].R = l
		array[i].S = s
		hitungLuasKelilingPersegi(s, &array[i].LP, &array[i].KP)
		hitungLuasKelilingLingkaran(l, &array[i].LL, &array[i].KL)
		hitungTotal(array[i].LL, array[i].LP, array[i].KL, array[i].KP, &array[i].TL, &array[i].TP)
		i++
		fmt.Scan(&l, &s)
	}
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	for j := 0; j < i; j++ {
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", array[j].R, array[j].S, array[j].LL, array[j].LP, array[j].KL, array[j].KP, array[j].TL, array[j].TP)
	}
}
func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	// r = radius , luas = l , keliling (1)
	*l = math.Pi * r * r
	*k = 2 * math.Pi * r
}
func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}
func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}
