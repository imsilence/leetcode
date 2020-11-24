package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}

EOF:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue EOF
			}
		}
		return i

	}
	return -1
}
