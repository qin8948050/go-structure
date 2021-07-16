/* 
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

 

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
] */

type TreeNode struct {
Val int
 Left *TreeNode
  Right *TreeNode
}
func LevelOrder(root *TreeNode) [][]int {
	res:=[][]int{}
	if root==nil{
		return res
	}
	q:=[]*TreeNode{root}
	index:=0
	subq:=[]*TreeNode{}
	res=append(res,[]int{})
	for len(q)>0 {
		if q[0].Left!=nil{
			subq=append(subq, q[0].Left)
		}
		if q[0].Right!=nil {
			subq=append(subq,q[0].Right)
		}
		res[index]=append(res[index],q[0].Val)
		q=q[1:]
		if len(q)==0 && len(subq)>0{
			q=subq
			index++
			subq=[]*TreeNode{}
			res=append(res,[]int{})
		}
	}
	return res
}