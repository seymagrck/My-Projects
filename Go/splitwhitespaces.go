package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	currentWord := ""
	inWord := false

	for _, char := range str {
		if char == ' ' || char == '\t' || char == '\n' {
			if inWord {
				result = append(result, currentWord)
				currentWord = ""
			}
			inWord = false
		} else {
			currentWord += string(char)
			inWord = true
		}
	}

	if inWord {
		result = append(result, currentWord)
	}

	return result
}
