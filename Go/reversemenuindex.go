package piscine

func ReverseMenuIndex(menu []string) []string {
	dizi2 := []string{}

	for i := len(menu) - 1; i >= 0; i-- {
		dizi2 = append(dizi2, menu[i])
	}

	return dizi2
}
