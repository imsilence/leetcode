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

// Int2Equals 验证二维整数切片元素是否相同
func Int2Equals(a, b [][]int, sorted bool) bool {
	if len(a) != len(b) {
		return false
	}
	if sorted {
		for _, l := range a {
			sort.Ints(l)
		}
		for _, l := range b {
			sort.Ints(l)
		}
		ints2 := func(list [][]int) {
			sort.Slice(list, func(i, j int) bool {
				for x, v := range list[i] {
					if v < list[j][x] {
						return true
					} else if v > list[j][x] {
						return false
					}
				}
				return false
			})
		}
		ints2(a)
		ints2(b)
	}
	for i, v := range a {
		if !IntEquals(v, b[i], false) {
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

// StringContains 验证字符串在字符串切片中
func StringContains(list []string, txt string) bool {
	for _, e := range list {
		if e == txt {
			return true
		}
	}
	return false
}
