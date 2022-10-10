/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    p1, p2 := head, head
    for k > 0 && p1 != nil {
        p1, k = p1.Next, k-1
    }
    
    for p1 != nil {
        p1, p2 = p1.Next, p2.Next
    }

    return p2
}
