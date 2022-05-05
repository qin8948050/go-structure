/*
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	var maxLen = 1
	var left, right, window = 0, 0, make(map[byte]bool)
	for right < n {
		rightChar := s[right]
		for window[rightChar] { //清空之前保留的字串
			delete(window, s[left])
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		window[rightChar] = true
		right++
	}
	return maxLen
}
