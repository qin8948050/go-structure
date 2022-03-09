/* 
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

 

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false */

/* 实现思路：
对称二叉树定义： 对于树中 任意两个对称节点 LL 和 RR ，一定有：
L.val = R.val L.val=R.val ：即此两对称节点值相等。
L.left.val = R.right.val L.left.val=R.right.val ：即 LL 的 左子节点 和 RR 的 右子节点 对称；
L.right.val = R.left.val L.right.val=R.left.val ：即 LL 的 右子节点 和 RR 的 左子节点 对称。
根据以上规律，考虑从顶至底递归，判断每对节点是否对称，从而判断树是否为对称二叉树。
算法流程：
isSymmetric(root) ：

特例处理： 若根节点 root 为空，则直接返回true。
返回值： 即 recur(root.left, root.right) ;
recur(L, R) ：

终止条件：
当 LL 和 RR 同时越过叶节点： 此树从顶至底的节点都对称，因此返回 true；
当 LL 或 RR 中只有一个越过叶节点： 此树不对称，因此返回 false ；
当节点 LL 值 = 节点 RR 值： 此树不对称，因此返回 falsefalse ；
递推工作：
判断两节点 L.left 和 R.right 是否对称，即 recur(L.left, R.right) ；
判断两节点 L.right 和 R.left 是否对称，即 recur(L.right, R.left) ；
参考： https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/solution/mian-shi-ti-28-dui-cheng-de-er-cha-shu-di-gui-qing/
 */

func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return recur(root.Left,root.Right)
}

func recur(left *TreeNode,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left ==nil || right==nil || left.Val!=right.Val{
		return false
	}
	return recur(left.Left,right.Right) && recur(left.Right,right.Left)
}