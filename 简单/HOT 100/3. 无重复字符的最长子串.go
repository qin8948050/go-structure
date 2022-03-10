func lengthOfLongestSubstring(s string) int {
	var n=len(s)
	if n<=1 {
		return n
	}
	var maxLen=1
	var left,right,window=0,0,make(map[byte]bool)
	for right<n {
		rightChar:=s[right]
		for window[rightChar] {  //清空之前保留的字串
			delete(window,s[left])
			left++
		}
		if right-left+1>maxLen {
			maxLen=right-left+1
		}
		window[rightChar]=true
		right++
	}
	return maxLen
}