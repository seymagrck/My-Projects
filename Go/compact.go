package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	compactSize := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[compactSize] = slice[i]
			compactSize++
		}
	}

	*ptr = slice[:compactSize]
	return compactSize
}
