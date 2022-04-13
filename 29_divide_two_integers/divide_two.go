package dividetwointegers

func divide(dividend int, divisor int) int {
	if dividend/divisor >= 2147483648 {
		return 2147483647
	}
	if dividend/divisor <= -2147483648 {
		return -2147483648
	}

	s := abs(divisor)

	c := 0
	for s <= abs(dividend) {
		c++
		s += abs(divisor)
	}
	if dividend < 0 && divisor > 0 {
		return -c
	} else if dividend > 0 && divisor < 0 {
		return -c
	}

	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
