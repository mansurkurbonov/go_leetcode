package plusone

import (
	"strconv"
	"testing"
)

func TestPlusOne(t *testing.T) {
	in := []struct {
		nums   []int
		expect []int
	}{
		{nums: []int{1, 2, 3}, expect: []int{1, 2, 4}},
		{nums: []int{4, 3, 2, 1}, expect: []int{4, 3, 2, 2}},
		{nums: []int{9, 9}, expect: []int{1, 0, 0}},
		{nums: []int{9}, expect: []int{1, 0}},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := plusOne(in[i].nums)
			for j := 0; j < len(output); j++ {
				if output[j] != in[i].expect[j] {
					t.Error("got", output, "expect", in[i].expect)
				}
			}
		})
	}
}
