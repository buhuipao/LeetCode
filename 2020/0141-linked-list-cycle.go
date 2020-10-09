/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast, slow = fast.Next.Next, slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}
