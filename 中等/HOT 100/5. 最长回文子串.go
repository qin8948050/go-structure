/* 给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a" */

//方法1：暴力解法
func longestPalindrome(s string) string {
	var maxSize int
	length := len(s)
	//记录子串的起始位置
	var begin int
	if length <= 2 {
		return string(s[0])
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxSize && Parse(s, i, j) {
				maxSize = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxSize]
}

func Parse(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//中心扩散法
func longestPalindrome(s string) string {
	length := len(s)
	var begin int
	var end int
	var maxSize int
	if length < 2 {
		return string(s[0])
	}
	for i := 0; i < length; i++ {
		//奇数匹配
		single := Parse(s, i, i)
		//偶数匹配
		double := Parse(s, i, i+1)
		maxSize = maxValue(single, double)
		if maxSize > end-begin+1 {
			//起始位置
			begin = i - (maxSize-1)/2
			end = i + maxSize/2
		}
	}
	return s[begin : end+1]
}

func Parse(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func maxValue(i, j int) int {
	if i > j {
		return i
	}
	return j
}