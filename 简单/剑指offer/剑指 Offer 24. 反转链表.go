/*

定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL 
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	pre:=new(ListNode)
	pre.Next=nil
	curr:=head
	for curr!= nil{
		p:=curr
		p.Next=pre.Next
		pre.Next=p
		curr=curr.Next
	}
	return pre.Next
}