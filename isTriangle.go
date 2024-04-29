package main

import (
    "fmt"
)

func isTriangle(a, b, c int) bool {
    // In a triangle, the sum of the lengths of any two sides must be greater than the length of the third side
    return (a+b > c) && (b+c > a) && (a+c > b)
}

func main() {
    var side1, side2, side3 int

    fmt.Println("Masukkan panjang sisi-sisi bidang:")
    fmt.Print("Sisi 1: ")
    fmt.Scanln(&side1)
    fmt.Print("Sisi 2: ")
    fmt.Scanln(&side2)
    fmt.Print("Sisi 3: ")
    fmt.Scanln(&side3)

 
    if isTriangle(side1, side2, side3) {
        fmt.Println("Segitiga")
    } else {
        fmt.Println("Bukan segitiga")
    }
}
