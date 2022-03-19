package searchinsertposition

import (
	"strconv"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	in := []struct {
		nums   []int
		target int
		expect int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expect: 4},
		{nums: []int{1, 3}, target: 3, expect: 1},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := searchInsert(in[i].nums, in[i].target)
			if output != in[i].expect {
				t.Error("got", output, "expect", in[i].expect)
			}
		})
	}
}
