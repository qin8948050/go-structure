
/*

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

 

示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
示例 4：

输入：x = -101
输出：false 
 */

 //1.最优：用到字符串翻转
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	res := 0
	tmp := 0
	maxLastNumber := math.MaxInt32 % 10
	minLastNumber := math.MinInt32 % 10
	maxNumber := math.MaxInt32 / 10
	minNumber := math.MinInt32 / 10
	for (x != 0) {
		tmp = x % 10
		if res > maxNumber || (res == maxNumber && tmp > maxLastNumber) {
			return 0
		}
		if res < minNumber || (res == minNumber && tmp < minLastNumber) {
			return 0
		}
		res = res*10 + tmp
		x = x / 10
	}
	return res
}
func isPalindrome(x int) bool {
	if x<0 {
		return false
	}
	res:=reverse(x)
	if res==x {
		return  true
	}
	return false
}

//2.转化为字符串，数组
func isPalindrome(x int) bool {
	i:=0
	j:=0
	if x<0 {
		return false
	}
	s:=strconv.Itoa(x)
	aa:=strings.Split(s,"")

	if len(s) == 1{
		return true
	}
	if len(aa)  == 2{
		if aa[0] == aa[1]{
			return true
		} else {
			return false
		}
	}
	for ;i<len(aa)&&j<len(aa)-1;i++{
		j+=2
	}
	med:=i
	i=0
	j=len(aa)-1
	for ; i<=med && j>=med;i++{
		if (aa[i] == aa[j]){
			j--
		} else {
			return false
		}
	}
	return true
}