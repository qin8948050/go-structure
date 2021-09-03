func height(root *TreeNode) float64 {
	if root==nil{
		return 0
	}
	l:=height(root.Left)
	r:=height(root.Right)
	if l<r {
		return r+1
	} else {
		return l+1
	}
}