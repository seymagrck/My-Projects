package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return false
		}
	}
	return true
}
