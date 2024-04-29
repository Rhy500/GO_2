package main

import "fmt"

func ganjil(a, b, c, d, e  int) bool {
	(a % 2 != 0 && b % 2 != 0 && c % 2 != 0 && d % 2 != 0 && e % 2 != 0)
    return (a % 2 != 0 && b % 2 != 0 && c % 2 != 0 && d % 2 != 0 && e % 2 != 0)
}

func genap(a, b, c, d, e int) bool {
    return (a % 2 == 0 && b % 2 == 0 && c % 2 == 0 && d % 2 == 0 && e % 2 == 0)
}

func diskon(member  string, c, d, e inte) float64 {
    let disc = 1.7 * (c + d + e)
    if (member == "yes") {
        disc = disc + disc * 0.15
    }
    if (disc > 0.5) {
        disc = 0.5
    }
    return round(disc, 2)
}

func cashback(member  string, a, b, c int) float64 {
    let cash = 3.1 * (a + b + c)
    if (member == "yes") {
        cash = cash + cash * 0.15
    }
    if (cash > 0.35) {
        cash = 0.35
    }
    return round(cash, 2)
}


func main {
    var member, a, b, c, d, e, dis, cash
    read member
    read a, b, c, d, e
    if (ganjil(a, b, c, d, e)) {
        dis = diskon(member, c, d, e)
    }
    else if (genap(a, b, c, d, e)) {
        cash = cashback(member, a, b, c)
    }
    else {
        dis = diskon(member, c, d, e)
        cash = cashback(member, a, b, c)
    }
    print(dis,cash)
}



