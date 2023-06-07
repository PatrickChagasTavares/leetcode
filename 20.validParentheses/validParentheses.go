package validparentheses

func isValid(s string) bool {
	matchies := make(map[rune]bool)
	caracterMap := map[rune]rune{
		'(': ')',
		')': '(',
		'[': ']',
		']': '[',
		'{': '}',
		'}': '{',
	}

	for _, char := range s {
		item := caracterMap[char]
		value, ok := matchies[char]

		if !ok {
			matchies[char] = !ok
			matchies[item] = value
			continue
		}
		matchies[char] = !value
	}

	for _, match := range matchies {
		if !match {
			return false
		}
	}

	return true
}

func isValidV2(s string) bool {
	caracterMap := map[rune]rune{
		'(': ')',
		')': '(',
		'[': ']',
		']': '[',
		'{': '}',
		'}': '{',
	}

	for i := 0; i < len(s)-1; i++ {
		curChar := rune(s[i])
		NextChar := rune(s[i+1])
		value := caracterMap[curChar]

		if i > 0 {
			if value != NextChar && value != rune(s[i-1]) {
				return false
			}
			continue
		}
		if value != NextChar {
			return false
		}
	}

	return true
}

func isValidV3(s string) bool {
	if s == "" {
		return true
	}

	matchies := []rune{}
	caracterMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		if _, ok := caracterMap[char]; ok {
			matchies = append(matchies, char)
			continue
		}

		value := caracterMap[matchies[len(matchies)-1]]
		if len(matchies) == 0 || value != char {
			return false
		}

		matchies = matchies[:len(matchies)-1]
	}

	return len(matchies) == 0
}
