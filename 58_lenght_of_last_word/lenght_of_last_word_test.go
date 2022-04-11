package lenghtoflastword

import (
	"strconv"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	in := []struct {
		in     string
		expect int
	}{
		{in: "Hello World", expect: 5},
		{in: "   fly me   to   the moon   ", expect: 4},
		{in: "luffy is still joyboy", expect: 6},
		{in: "liverpool is a best football club in the world   !", expect: 1},
	}

	for i := 0; i < len(in); i++ {
		t.Run("test № "+strconv.Itoa(i+1), func(t *testing.T) {
			output := lengthOfLastWord(in[i].in)
			if output != in[i].expect {
				t.Error("got", output, "expect", in[i].expect)
			}
		})
	}
}
