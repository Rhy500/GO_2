package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += g - k
}

func hitungJumPoint(jp *int) {
	*jp = jm*3 + jd
}

func main() {
	var n int
	fmt.Print("Masukan banyaknya pertandingan (N): ")
	fmt.Scan(&n)

	var jm, jd, jk, jg, jk, jsg, jp int
	for i := 0; i < n; i++ {
		var g, k int
		var hasil string
		fmt.Print("Masukan hasil pertandingan (" + strconv.Itoa(i+1) + "): ")
		fmt.Scan(&hasil)
		arrHasil := strings.Fields(hasil)
		g, _ = strconv.Atoi(arrHasil[0])
	}
}
