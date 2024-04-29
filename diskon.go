package main

import "fmt"

func main() {
	var total float64
	var member bool
	fmt.Scan(&total, &member)
	fmt.Println(diskonBelanja(total, member))
}
func diskonBelanja(total float64, member bool) float64 {
	var diskon float64
	//if member == true && total > 100000 {
	//return (total - (total * (10 / 100)))
	//} else {
	//return (total - (total * (5 / 100)))
	//}
	if total > 100000 {
		if member {
			diskon = 0.1
		} else {
			diskon = 0.05
		}
	} else {
		diskon = 0
	}
	return total * (1 - diskon)
}
