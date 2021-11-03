package piscine

func DivMod(a int, b int, div *int, mod *int) {
	a = 13
	b = 2

	*div = a / b

	*mod = a % b
}
