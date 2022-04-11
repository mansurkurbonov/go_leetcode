package lenghtoflastword

func lengthOfLastWord(s string) int {
	words := []string{}

	word := ""
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 && string(s[i]) != " " {
			word = word + string(s[i])
			words = append(words, word)
			continue
		}
		if string(s[i]) != " " {
			word = word + string(s[i])
		} else if string(s[i]) == " " {
			if word != "" {
				words = append(words, word)
			}
			word = ""
		}
	}

	return len(words[len(words)-1])
}
