package main

import "fmt"

func main() {
	recurse_master()
}

func recurse_master() {
	// IS: -
	// Proses: Membaca karakter, menginisialisasi variabel jumlah,
	//         memanggil prosedur recurse, dan mencetak jumlah
	// FS: Tercetak jumlah vokal lowercase di layar
	var jumlah int
	recurse(' ', &jumlah)
	fmt.Println(jumlah)
}

func recurse(kar byte, jumlah *int) {
	// IS: kar dan jumlah terdefinisi
	// Proses: Jika terbaca '.' mengupdate nilai variabel jumlah hanya jika kar
	//         adalah karakter vokal. Jika terbaca bukan karakter '.', maka
	//         akan mengupdate nilai variabel jumlah hanya jika kar adalah vokal,
	//         membaca kar, dan memanggil prosedur recurse. Panggil fungsi vowel
	//         untuk memeriksa apakah kar termasuk vokal lowercase atau tidak.
	// FS: jumlah berisi nilai
	if kar == '.' {
		return
	} else {
		if vowel(kar) {
			*jumlah++
		}
		fmt.Scanf("%c", &kar)
		recurse(kar, jumlah)
	}
}

func vowel(kar byte) bool {
	// mengembalikan boolean true jika kar adalah huruf vokal lowercase
	switch kar {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
