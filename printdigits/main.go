package main

import "github.com/01-edu/z01"

func main() {
	num := "0123456789"

	for i := 0; i <= 9; i++ {
		z01.PrintRune(rune(num[i]))
	}
	z01.PrintRune(rune('\n'))
}
