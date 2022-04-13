package medianoftwosortedarrays

import (
	"strconv"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	in := []struct {
		nums1  []int
		nums2  []int
		expect float64
	}{
		{nums1: []int{1, 3}, nums2: []int{2}, expect: 2},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, expect: 2.5},
		{nums1: []int{1, 2, 5}, nums2: []int{3, 4, 6}, expect: 3.5},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := findMedianSortedArrays(in[i].nums1, in[i].nums2)
			if output != in[i].expect {
				t.Error("got", output, "expect", in[i].expect)
			}
		})
	}
}
