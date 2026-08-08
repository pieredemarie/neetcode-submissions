/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
 
public class Solution {
    public ListNode MergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 is null)
            return list2;
        if (list2 is null)
            return list1;
        ListNode nHead = null;
        if (list1.val <= list2.val)
        {
            nHead = list1;
            list1 = list1.next;
        } 
        else
        {
            nHead = list2;
            list2 = list2.next;
        }

        ListNode curr = nHead;
        while (list1 is not null && list2 is not null)
        {
            if (list1.val >= list2.val)
            {
                curr.next = list2;
                list2 = list2.next;
            }
            else 
            {
                curr.next = list1;
                list1 = list1.next;
            }
            curr = curr.next;
        }
        if (list1 is not null)
            curr.next = list1;
        else
            curr.next = list2;
        return nHead;
    }
}