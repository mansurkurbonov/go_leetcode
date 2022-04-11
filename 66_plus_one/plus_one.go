package plusone

func plusOne(digits []int) []int {
	res := []int{}

	plus := 1
	inMind := 0

	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + plus + inMind
		if sum > 9 {
			a := sum % 10
			inMind = sum / 10
			res = append(res, a)
			plus = 0
		} else {
			res = append(res, sum)
			inMind = 0
			plus = 0
		}
		if i == 0 && inMind > 0 {
			res = append(res, inMind)
		}

	}

	revers := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		revers = append(revers, res[i])
	}

	return revers
}
