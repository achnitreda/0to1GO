package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a % *b
	*a = *a / *b
	*b = temp
}
