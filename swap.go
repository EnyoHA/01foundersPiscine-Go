package piscine

func Swap(a *int, b *int) {
	var aSwap int
	var bSwap int

	aSwap = *a
	bSwap = *b

	*a = bSwap
	*b = aSwap
}
