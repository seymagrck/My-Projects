package piscine

func Map(f func(int) bool, a []int) []bool {
	isprime := make([]bool, len(a))
	for i, n := range a {
		isprime[i] = f(n)
	}
	return isprime
}
