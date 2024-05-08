package piscine

func BasicAtoi2(s string) int {
	var result int
	multiplier := 1
	valid := true

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0') // Karakteri rakama dönüştür
			result += digit * multiplier
			multiplier *= 10
		} else if s[i] != ' ' {
			valid = false
			break
		}
	}

	if !valid {
		return 0
	}

	return result
}
