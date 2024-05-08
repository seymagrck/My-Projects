package piscine

func DescendAppendRange(max, min int) []int {
	var arr []int
	if max > min {
		for i := max; i > min; i-- {
			arr = append(arr, i)
		}
	} else {
		return []int{}
	}
	return arr
}
