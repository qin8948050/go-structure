/* 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true */

func isValid(s string) bool {
	map1 := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i, v1 := range s {
		if v, ok := map1[byte(v1)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}