package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	in := []struct {
		nums []int
		val  int
		out  int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, out: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, out: 5},
	}
	t.Run("1 test", func(t *testing.T) {
		output := removeElement(in[0].nums, in[0].val)
		if output != in[0].out {
			t.Error("got", output, "expect", in[0].out)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := removeElement(in[1].nums, in[1].val)
		if output != in[1].out {
			t.Error("got", output, "expect", in[1].out)
		}
	})
}
