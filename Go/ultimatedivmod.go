package piscine

func UltimateDivMod(a *int, b *int) {
	if b != nil && a != nil && *b != 0 {
		temp := *a
		*a = temp / *b

		*b = temp % *b
	}
}
