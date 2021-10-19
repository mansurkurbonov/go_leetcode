package romantointeger

import (
	"fmt"
	"strings"
)

func romanToInt(rom string) int {
	if checkRoman(rom) {
		return 0
	}
	s := strings.Split(rom, "")

	sum := 0
	for i := 0; i < len(s); i++ {
		_, ok := romanMap[string(s[i])]
		if !ok {
			return 0
		}

		if i != len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}

		fmt.Println(sum)
	}

	return sum
}

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func checkRoman(s string) bool {
	return len(s) < 1 || len(s) > 15
}
