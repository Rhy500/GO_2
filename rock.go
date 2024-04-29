package main

import (
	"fmt"
)

func pukulan(P, X, Y, Z, H int) (total int) {
	for i := 1; i <= H; i++ {
		if i%X == 0 || i%Y == 0 {
			if i%Z != 0 {
				total += P
			}
		}
	}
	return
}

func tendangan(T, X, Y, Z, H int) (total int) {
	for i := 1; i <= H; i++ {
		if i%X == 0 && i%Y == 0 {
			if i%Z != 0 {
				total += T
			}
		}
	}
	return
}

func delapan_gerbang(G, X, Y, Z, H int) (total int) {
	for i := 1; i <= H; i++ {
		if i%X == 0 && i%Y != 0 {
			if i%Z != 0 {
				total += G
			}
		}
	}
	return
}

func main() {
	var P, T, G, X, Y, Z, H int
	fmt.Scan(&P, &T, &G, &X, &Y, &Z, &H)
	fmt.Printf("%d %d %d\n", pukulan(P, X, Y, Z, H), tendangan(T, X, Y, Z, H), delapan_gerbang(G, X, Y, Z, H))
}
