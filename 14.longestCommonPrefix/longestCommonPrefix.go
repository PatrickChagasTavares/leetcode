package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	lenght := len(strs)
	prefix := strs[0]

	for _, s := range strs {
		i := 0
		for ; i < lenght && i < len(prefix) && prefix[i] == s[i]; i++ {
		}
		prefix = prefix[:i]

	}
	return prefix
}

func longestCommonPrefixV2(strs []string) string {
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
