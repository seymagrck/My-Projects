package piscine

func Atoi(s string) int {
	var result int
	multiplier := 1
	valid := true
	sign := 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			result += digit * multiplier
			multiplier *= 10
		} else if s[i] == '-' && i == 0 {
			sign = -1
		} else if s[i] == '+' && i == 0 {
			sign = 1
		} else if s[i] != ' ' {
			valid = false
			break
		}
	}

	if !valid || (result == 0 && multiplier != 1) {
		return 0
	}

	return result * sign
}
