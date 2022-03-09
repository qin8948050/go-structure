/* 请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

 

示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1] */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirrorTree(root *TreeNode) *TreeNode {
	if root==nil {
		return nil
	}
	t:=root
	vars:=[]*TreeNode{root}
	vars[0]=root
	for len(vars)>0{
		u:=vars[0]
		vars=vars[1:]
		u.Left,u.Right=u.Right,u.Left
		if u.Left !=nil {
			vars=append(vars,u.Left)
		}
		if u.Right!=nil {
			vars=append(vars,u.Right)
		}
	}
	return t
}