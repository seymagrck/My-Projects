package piscine

func ConcatParams(args []string) string {
	str := ""

	for i, rep := range args {
		str += string(rep)
		if i != len(args)-1 {
			str += "\n"
		}
	}
	return str
}
