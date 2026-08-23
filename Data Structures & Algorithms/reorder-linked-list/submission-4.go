/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
    fast, slow := head, head 
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr 
		curr = next
	}

	first := head
	second := prev 

	for second.Next != nil {
		nextFirst := first.Next
		nextSecond := second.Next
		
		first.Next = second
		second.Next = nextFirst

		first = nextFirst
		second = nextSecond
	}
}
