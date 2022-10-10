/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    // 环内相遇
    for fast != nil && fast.Next != nil {
        fast, slow = fast.Next.Next, slow.Next
        if fast == slow {
            break
        }
    }

    // 无环
    if fast == nil || fast.Next == nil {
        return nil
    }

    // 入环点相遇
    slow = head
    for fast != slow {
        fast, slow = fast.Next, slow.Next
    }

    return slow
}
