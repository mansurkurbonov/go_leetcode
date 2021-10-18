package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	in := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	t.Run("1 test", func(t *testing.T) {
		output := twoSum(in[0].nums, in[0].target)
		if output[0] != in[0].expected[0] && output[1] != in[0].expected[1] {
			t.Error("got", output, "expect", in[0].expected)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := twoSum(in[1].nums, in[1].target)
		if output[0] != in[1].expected[0] && output[1] != in[1].expected[1] {
			t.Error("got", output, "expect", in[1].expected)
		}
	})

	t.Run("3 test", func(t *testing.T) {
		output := twoSum(in[2].nums, in[2].target)
		if output[0] != in[2].expected[0] && output[1] != in[2].expected[1] {
			t.Error("got", output, "expect", in[2].expected)
		}
	})

}
