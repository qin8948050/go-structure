/* 给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0 
         / \ 
       -3   9 
       /   / 
     -10  5  */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	if len(nums)==1 {
		return &TreeNode{Val:nums[0]}
	}
	middle:=len(nums)/2
	head := &TreeNode{Val:nums[middle]}
	head.Left=sortedArrayToBST(nums[0:middle])
	head.Right=sortedArrayToBST(nums[middle+1:])
	return head
}