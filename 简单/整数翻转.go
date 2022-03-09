/* 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
 

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0 */

func reverse(x int) int {
	if x==0 {
		return 0
	}
	res:=0
	tmp:=0
	maxLastNumber:=math.MaxInt32%10
	minLastNumber:=math.MinInt32%10
	maxNumber:=math.MaxInt32/10
	minNumber:=math.MinInt32/10
	for(x!=0){
		tmp=x%10
		if res>maxNumber || (res==maxNumber && tmp>maxLastNumber) {
			return 0
		}
		if res< minNumber || (res==minNumber && tmp<minLastNumber) {
			return 0
		}
		res=res*10+tmp
		x=x/10
	}
	return res
}