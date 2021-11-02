package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb <= -1 {
		z01.PrintRune(rune('T'))
	} else {
		z01.PrintRune(rune('F'))
	}
	z01.PrintRune(rune('\n'))
}
