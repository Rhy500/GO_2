package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	display(a, b)
}

func isFaktor(num, x int) bool {
	/* mengembalikan true apabila x adalah faktor dari num */
	return num%x == 0
}

func jumlahFaktor(num int, total *int) {
	/* I.S. terdefinsi bilangan bulat num
	   F.S. total berisi hasil penjumlahan semua faktor dari num (tidak termasuk num) */
	*total = 0
	for i := 1; i <= num/2; i++ {
		if isFaktor(num, i) {
			*total += i
		}
	}
}

func perfect(num int) bool {
	/* mengembalikan true apabila num adalah perfect number, panggil jumlahFaktor() di sini*/
	var total int
	jumlahFaktor(num, &total)
	return total == num
}

func display(x, y int) {
	/* I.S. terdefinsi bilangan bulat x dan y
	   F.S. menampilkan barisan perfect number dari x hingga y */
	for i := x; i <= y; i++ {
		if perfect(i) {
			fmt.Print("Barisan Perfect Number: ", i, " ")
		}
	}
}
