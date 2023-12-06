package mathutils

func SumSlice[S ~[]E, E any, N int|float32|float64](s S, f func(E) N) N {
	var sum N = 0
	for _, v := range s {
		val := f(v)
		sum += val
	}
	return sum
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
