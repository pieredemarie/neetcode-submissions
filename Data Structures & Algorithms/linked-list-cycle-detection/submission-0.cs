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
    public bool HasCycle(ListNode head) {
        HashSet<ListNode> hs = new HashSet<ListNode>();
        ListNode t = head;

        while (t is not null)
        {
            if (hs.Contains(t))
                return true;
            hs.Add(t);
            t = t.next;
        }   
        return false;
    }
}
