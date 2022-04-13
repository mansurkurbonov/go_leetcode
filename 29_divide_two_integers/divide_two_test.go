package dividetwointegers

import (
	"strconv"
	"testing"
)

func TestDivide(t *testing.T) {
	in := []struct {
		dividend int
		divisor  int
		expect   int
	}{
		{dividend: 10, divisor: 3, expect: 3},
		{dividend: 7, divisor: -3, expect: -2},
		{dividend: 10, divisor: 2, expect: 5},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := divide(in[i].dividend, in[i].divisor)
			if output != in[i].expect {
				t.Error("got", output, "expect", in[i].expect)
			}
		})
	}
}
