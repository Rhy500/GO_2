package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var pilih int
	intro()
	for {
		menu_utama(&pilih)
		switch pilih {
		case 1:
			jumlah_bilangan_asli()
		case 2:
			faktorial()
		case 3:
			fibonacci()
		default:
			clear_screen()
		}
		if pilih == 4 {
			break
		}
	}
	bye()
}

func intro() {
	clear_screen()
	fmt.Println("Selamat datang")
}

func bye() {
	clear_screen()
	fmt.Println("Sampai Jumpa")
}
func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
func menu_utama(p *int) {
	fmt.Println("---------------------------")
	fmt.Println("			MENU			")
	fmt.Println("---------------------------")
	fmt.Println("1. Jumlah Bilangan Asli 	")
	fmt.Println("2. Faktorial				")
	fmt.Println("3. Fibonacci				")
	fmt.Println("4. Exit					")
	fmt.Println("---------------------------")
	fmt.Println("pilih (1/2/3/4):			")
	fmt.Scan(p)
}
func menu_jumlah_bilangan_asli() {
	clear_screen()
	fmt.Println("---------------------------")
	fmt.Println("			MENU			")
	fmt.Println("---------------------------")
	fmt.Println("1. Jumlah Bilangan Asli 	")
	fmt.Println("2. Exit					")
	fmt.Println("---------------------------")
}
func faktorial() {
	var pilih, N, suku int
	var lagi string
	for {
		menu_faktorial()
		fmt.Print("pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			for {
				fmt.Print("masukan bilan asli")
				fmt.Scan()
			}
		}
	}
}
