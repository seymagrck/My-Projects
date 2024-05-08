package piscine

func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
}
