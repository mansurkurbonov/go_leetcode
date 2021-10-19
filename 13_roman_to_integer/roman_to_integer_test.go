package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	in := []struct {
		in  string
		out int
	}{
		{in: "III", out: 3},
		{in: "IV", out: 4},
		{in: "LVIII", out: 58},
		{in: "MCMXCIV", out: 1994},
	}

	t.Run("1 test", func(t *testing.T) {
		output := romanToInt(in[0].in)
		if output != in[0].out {
			t.Error("got", output, "expect", in[0].out)
		}
	})

	t.Run("2 test", func(t *testing.T) {
		output := romanToInt(in[1].in)
		if output != in[1].out {
			t.Error("got", output, "expect", in[1].out)
		}
	})

	t.Run("3 test", func(t *testing.T) {
		output := romanToInt(in[2].in)
		if output != in[2].out {
			t.Error("got", output, "expect", in[2].out)
		}
	})

	t.Run("4 test", func(t *testing.T) {
		output := romanToInt(in[3].in)
		if output != in[3].out {
			t.Error("got", output, "expect", in[3].out)
		}
	})

}
