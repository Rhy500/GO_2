package main

import "fmt"

func main() {
	var lumin int

	fmt.Scan(&lumin)

	fmt.Println(Quantum_Shard, Galactum, Nebula_Crown, Stellaris, Lumin)
}

func Quantum_Shard(lumin int) {
	quantumShard := lumin / 10 / 3 / 2 / 20
	lumin = lumin - quantumShard*10*3*2*20
}
func Galactum(lumin int) {
	galactum := lumin / 3 / 2 / 20
	lumin = lumin - galactum*3*2*20
}
func Nebula_Crown(lumin int) {
	nebulaCrown := lumin / 2 / 20
	lumin = lumin - nebulaCrown*2*20
}
func Stellaris(lumin int) {
	stellaris := lumin / 20
	lumin = lumin - stellaris*20
}
func Lumin(lumin int) {
	lumin = lumin - lumin
}
