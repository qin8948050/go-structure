
/* 给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。

 

示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等） */
/*
解题思路：
借助map来计数，基数个数为2的倍数或为1；
定义一个全局变量n,来记录基数个数不为2的倍数的情况。 
 */
func canPermutePalindrome(s string) bool {
	hash := make(map[int32]int)
	n := 0
	for _,v := range s{
		hash[v]++
	}
	for _,v:=range hash{
		if v%2!=0{  
			n++
			if n==2{
				return false
			}
		}
	}
	return true
}