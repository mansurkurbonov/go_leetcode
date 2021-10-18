package twosum

import "math"

func twoSum(nums []int, target int) []int {
	if checkArray(nums) || checkNumber(target) {
		return []int{-1, -1}
	}

	m := map[int]int{}
	for i, v := range nums {
		if checkNumber(v) {
			return []int{-1, -1}
		}
		index2, ok := m[target-v]
		if ok {
			return []int{index2, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}

func checkArray(nums []int) bool {
	return len(nums) < 2 || len(nums) > 1000
}

func checkNumber(number int) bool {
	return float64(number) <= math.Pow(-10, 9) || float64(number) >= math.Pow(10, 9)
}
