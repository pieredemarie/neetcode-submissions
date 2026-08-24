/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{
		Next: head,
	}
	pre := dummy

	for i := 0;i<left-1;i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start.Next

	for i := 0;i<right-left;i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then 
		then = start.Next
	}

	return dummy.Next
}
