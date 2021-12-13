package intmath

func Min(i int, is ...int) int {
	min := i
	for _, i2 := range is {
		if i2 < min {
			min = i2
		}
	}
	return min
}

func Max(i int, is ...int) int {
	max := i
	for _, i2 := range is {
		if i2 > max {
			max = i2
		}
	}
	return max
}
