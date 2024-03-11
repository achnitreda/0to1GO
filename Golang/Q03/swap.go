package piscine

func Swap(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}
