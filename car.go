package main

import (
	"fmt"
	"strings"
)

type Character byte

const (
	Path       Character = '.'
	Orc        Character = 'O'
	Smaug      Character = 'S'
	Harta      Character = 'H'
	Arkenstone Character = 'A'
	End        Character = 'E'
)

type Position struct {
	bilbo, kurcaci, orc, smaug, harta, arkenstone, pos int
}

func (p *Position) step(c Character) {
	switch c {
	case Orc:
		p.orc--
		if p.orc <= 0 {
			p.bilbo = 0
			p.kurcaci = 0
		}
	case Smaug:
		p.smaug--
		if p.smaug <= 0 {
			p.bilbo = 0
			p.kurcaci = 0
		}
	case Harta:
		p.harta++
		p.bilbo = 100
	case Arkenstone:
		p.arkenstone = 1
	case Path:
		p.pos++
		if p.bilbo < 100 {
			p.bilbo++
		}
		if p.kurcaci < 13 {
			p.kurcaci++
		}
	}
}

func main() {
	var input = []string{
		"100 20 30",
		"..0.0.HH.A.S.E",
	}
	var p Position
	fmt.Scanf("%d %d %d %d\n", &p.bilbo, &p.kurcaci, &p.orc, &p.smaug)
	r := strings.Fields(input[1])
	for _, c := range r {
		p.step(Character(c[0]))
	}
	fmt.Printf("%d %d %d %d %d %d %t\n", p.pos, p.bilbo, p.kurcaci, p.harta, p.arkenstone, p.arkenstone, p.arkenstone > 0)
}
