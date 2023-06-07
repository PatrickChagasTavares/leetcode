package main

func isSubsequence(s string, t string) bool {
	k := 0
	isSubsequence := false
	if (s == "" && t == "") || s == "" {
		return true
	}

	for _, v := range t {
		if v == rune(s[k]) {
			k++
		}

		if k >= len(s) {
			isSubsequence = true
			break
		}
	}

	return isSubsequence
	// var c, l = 0, len(s) - 1

	// fmt.Println(len(t))
	// for i := 0; i < len(t); i++ {
	// 	if c == len(s) {
	// 		return true
	// 	}

	// 	if c < l && s[c+1] == t[i] {
	// 		return false
	// 	}

	// 	if s[c] == t[i] {
	// 		c++
	// 	}
	// }

	// return c == len(s)
}
