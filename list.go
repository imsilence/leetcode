package leetcode

import "sort"

// IntEquals 验证整数切片元素是否相同
func IntEquals(a, b []int, sorted bool) bool {
	if len(a) != len(b) {
		return false
	}
	if sorted {
		sort.Ints(a)
		sort.Ints(b)
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// StringEquals 验证字符串切片元素是否相同
func StringEquals(a, b []string, sorted bool) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
