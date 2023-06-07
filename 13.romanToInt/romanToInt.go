package romantoint

func romanToInt(s string) int {
	romanMps := createDictionary()
	sum := 0

	for i := 0; i < len(s); i++ {
		var numRoman, numRomanNext byte
		numRoman = s[i]
		value, ok := romanMps[rune(numRoman)]
		if !ok {
			sum = 0
			continue
		}

		if (i + 1) < len(s) {
			numRomanNext = s[i+1]
		}

		if sub := requiredSubtraction(rune(numRoman), rune(numRomanNext)); sub {
			sum -= value
			continue
		}

		sum += value
	}
	return sum
}
func romanToIntV2(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := romanMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}
	}

	return sum
}

func romanToIntV3(s string) int {
	romanMps := createDictionary()
	length := len(s)
	sum, ok := romanMps[rune(s[length-1])]
	if !ok {
		sum = 0
	}

	for i := length - 2; i >= 0; i-- {
		numRoman := s[i]
		value, ok := romanMps[rune(numRoman)]
		if !ok {
			continue
		}

		if requiredSubtraction(rune(numRoman), rune(s[i+1])) {
			sum -= value
			continue
		}

		sum += value
	}
	return sum
}

func romanToIntV4(s string) int {
	romanMps := createDictionary()
	length := len(s)
	sum, ok := romanMps[rune(s[length-1])]
	if !ok {
		sum = 0
	}

	for i := length - 2; i >= 0; i-- {
		numRoman := s[i]
		value, ok := romanMps[rune(numRoman)]
		if !ok {
			continue
		}

		if rune(numRoman) < rune(s[i+1]) {
			sum -= value
			continue
		}

		sum += value
	}
	return sum
}

func createDictionary() map[rune]int {
	i, v, x, l, c, d, m := 'I', 'V', 'X', 'L', 'C', 'D', 'M'
	return map[rune]int{
		i: 1,
		v: 5,
		x: 10,
		l: 50,
		c: 100,
		d: 500,
		m: 1000,
	}
}

func createDictionaryV2() map[byte]int {
	return map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
}

func requiredSubtraction(value, nextValue rune) bool {
	switch value {
	case 'I':
		if nextValue == 'V' || nextValue == 'X' {
			return true
		}
	case 'X':
		if nextValue == 'L' || nextValue == 'C' {
			return true
		}
	case 'C':
		if nextValue == 'D' || nextValue == 'M' {
			return true
		}
	}

	return false
}
func requiredSubtractionV2(value, nextValue byte) bool {
	switch value {
	case 'I':
		if nextValue == 'V' || nextValue == 'X' {
			return true
		}
	case 'X':
		if nextValue == 'L' || nextValue == 'C' {
			return true
		}
	case 'C':
		if nextValue == 'D' || nextValue == 'M' {
			return true
		}
	}

	return false
}

func GoogleV2(s string) int {
	characterMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return characterMap[s[0]]
	}

	sum := characterMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if characterMap[s[i]] < characterMap[s[i+1]] {
			sum -= characterMap[s[i]]
		} else {
			sum += characterMap[s[i]]
		}
	}

	return sum
}
