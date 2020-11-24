package leetcode

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/

func clone(stacks []*Stack) []*Stack {
	nss := make([]*Stack, len(stacks))
	for i, stack := range stacks {
		nss[i] = stack.Clone()
	}
	return nss
}

func generate(stacks []*Stack, n int, start, end int) []*Stack {
	if start == n && end == n {
		return stacks
	}

	if start == end {
		// 只能开始
		for _, stack := range stacks {
			stack.Push("(")
		}
		stacks = generate(stacks, n, start+1, end)
	} else if start == n {
		// 只能结束
		for _, stack := range stacks {
			stack.Push(")")
		}
		stacks = generate(stacks, n, start, end+1)
	} else {
		// 开始
		nss := clone(stacks)
		for _, stack := range nss {
			stack.Push("(")
		}
		nss = generate(nss, n, start+1, end)

		// 结束
		for _, stack := range stacks {
			stack.Push(")")
		}
		stacks = generate(stacks, n, start, end+1)

		// 合并
		stacks = append(stacks, nss...)
	}

	return stacks
}

func generateParenthesis(n int) []string {
	start, end := 1, 0

	stack := NewStack()
	stack.Push("(")
	stacks := []*Stack{stack}

	stacks = generate(stacks, n, start, end)

	result := make([]string, len(stacks))
	for i, stack := range stacks {
		result[i] = stack.Elements()
	}

	return result
}
