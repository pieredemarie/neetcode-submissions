/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    l1 := list1 
	l2 := list2 

	dummy := &ListNode{Val: 0, Next: nil}
	curr := dummy

	for l1 != nil && l2 != nil {
		if (l1.Val >= l2.Val) {
			curr.Next = l2 
			l2 = l2.Next
		} else if l1.Val < l2.Val {
			curr.Next = l1 
			l1 = l1.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
