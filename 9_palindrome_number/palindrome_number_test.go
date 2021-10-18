package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	in := []struct {
		val int
		out bool
	}{
		{val: 121, out: true},
		{val: -121, out: false},
		{val: 10, out: false},
	}
	t.Run("1 test", func(t *testing.T) {
		output := isPalindrome(in[0].val)
		if output != in[0].out {
			t.Error("got", output, "expect", in[0].out)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := isPalindrome(in[1].val)
		if output != in[1].out {
			t.Error("got", output, "expect", in[1].out)
		}
	})
	t.Run("3 test", func(t *testing.T) {
		output := isPalindrome(in[2].val)
		if output != in[2].out {
			t.Error("got", output, "expect", in[2].out)
		}
	})

}
