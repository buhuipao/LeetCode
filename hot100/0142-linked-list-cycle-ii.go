/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head    
    var node *ListNode
    for fast != nil && fast.Next != nil {
        fast, slow = fast.Next.Next, slow.Next
        if fast == slow {
            node = fast
            break
        }
    }
    // 无环
    if node == nil {
        return nil
    }
    fast = head
    for fast != slow {
        fast, slow = fast.Next, slow.Next
    }
    return fast
}
