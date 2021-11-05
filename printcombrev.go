package piscine

import "github.com/01-edu/z01"

func PrintCombRev() {
	for num1 := 9; num1 >= 2; num1-- {
		for num2 := num1 - 1; num2 >= 1; num2-- {
			for num3 := num2 - 1; num3 >= 0; num3-- {
				z01.PrintRune(rune(num1 + 48))
				z01.PrintRune(rune(num2 + 48))
				z01.PrintRune(rune(num3 + 48))
				if num1 > 2 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}

func piscine() {
	PrintCombRev()
}
