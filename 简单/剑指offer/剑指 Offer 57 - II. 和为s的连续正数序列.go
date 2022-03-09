/* 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

 

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]] */

/* 解法：滑动窗口 */



func findContinuousSequence(target int) [][]int {
	if target <= 0 {
		return nil
	}
	i := 1
	j := 1
	sum := 0
	result := [][]int{}
	for i <= target/2 {
		if sum < target {
			sum = sum + j
			j++
		} else if sum > target {
			sum = sum - i
			i++
		} else {
				tmp := []int{}
				for k := i; k < j; k++ {
					tmp = append(tmp, k)
				}
				result = append(result, tmp)
				sum=sum-i
				i++
		}
	}
	return result
}