/* URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

 

示例 1：

输入："Mr John Smith    ", 13
输出："Mr%20John%20Smith"
示例 2：

输入："               ", 5
输出："%20%20%20%20%20" */
//注意： byte类型为uint8的别名

func replaceSpaces(S string, length int) string {
	var arr []byte
	for i := 0 ;i < length ; i++ {
		if S[i] == 32 {
			arr = append(arr,37,50,48)
		}else {
			arr = append(arr,S[i])
		}
	}
	return string(arr)
}