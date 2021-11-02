package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	a := '1'
	b := '2'
	for c := '0'; c <= '7'; c++ {
		for a = c + 1; a <= '8'; a++ {
			for b = a + 1; b <= '9'; b++ {
				z01.PrintRune(c)
				z01.PrintRune(a)
				z01.PrintRune(b)
				if c < '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
