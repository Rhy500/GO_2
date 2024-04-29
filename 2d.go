package main

import (
	"fmt"
)

func main() {
	var r, s, l, k float64

	for {
		fmt.Scan(&l, &s)

		if l == 0 && s == 0 {
			break
		}
	}
	//fmt.Prinf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
}
func hitungLuasKelilingLingkaran(r float64, i, k *float64) {
	// r = radius , luas = l , keliling (1)
	pi := 3.14
	l = pi * r * r
	k = 2 * pi * r
}
func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	l = s * s
	k = 4 * s
}
func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	totLuas = lL + lP
	totKel = kL + kP
}
