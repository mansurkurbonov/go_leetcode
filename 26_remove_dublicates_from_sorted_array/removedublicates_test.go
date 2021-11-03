package removedublicatesfromsortedarray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	in := []struct {
		in  []int
		out int
	}{
		{in: []int{1, 1, 2}, out: 2},
		{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, out: 5},
	}
	t.Run("1 test", func(t *testing.T) {
		output := removeDuplicates(in[0].in)
		if output != in[0].out {
			t.Error("got", output, "expect", in[0].out)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := removeDuplicates(in[1].in)
		if output != in[1].out {
			t.Error("got", output, "expect", in[1].out)
		}
	})
}
