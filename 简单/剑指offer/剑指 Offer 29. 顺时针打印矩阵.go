/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

 

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7] 
 */

 /*
 解题思路：
 根据题目示例 matrix = [[1,2,3],[4,5,6],[7,8,9]] 的对应输出 [1,2,3,6,9,8,7,4,5] 可以发现，顺时针打印矩阵的顺序是 “从左向右、从上向下、从右向左、从下向上” 循环
算法流程：
空值处理： 当 matrix 为空时，直接返回空列表 [] 即可。
初始化： 矩阵 左、右、上、下 四个边界 l , r , t , b ，用于打印的结果列表 res 。
循环打印： “从左向右、从上向下、从右向左、从下向上” 四个方向循环，每个方向打印中做以下三件事 （各方向的具体信息见下表） ；
根据边界打印，即将元素按顺序添加至列表 res 尾部；
边界向内收缩 11 （代表已被打印）；
判断是否打印完毕（边界是否相遇），若打印完毕则跳出。
  */

func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0{
		return nil
	}
	l:=0
	r:=len(matrix[0])-1
	t:=0
	b:=len(matrix)-1
	var res []int
	for {
		//从左到右
		for i:=l;i<=r;i++{
			res=append(res,matrix[l][i])
		}
		t++
		if t>b {
			break
		}
		//从上到下
		for i:=t;i<=b;i++{
			res=append(res,matrix[i][r])
		}
		r--
		if (r<l) {
			break
		}
		//从右往左
		for i:=r;i>=l;i--{
			res=append(res,matrix[b][i])
		}
		b--
		if (b<t) {
			break
		}
		//从下往上
		for i:=b;i>=t;i--{
			res=append(res,matrix[i][l])
		}
		l++
		if l>r{
			break
		}
	}
	return res
}
