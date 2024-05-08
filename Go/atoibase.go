package piscine

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	charSet := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}

	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}

	result := 0
	baseLen := len(base)

	for _, char := range s {
		if index, exists := baseMap[char]; exists && index < baseLen {
			result = result*baseLen + index
		} else {
			return 0
		}
	}

	return result
}
