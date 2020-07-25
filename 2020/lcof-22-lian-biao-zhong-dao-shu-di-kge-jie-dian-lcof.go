/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针法
func getKthFromEnd(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    l, r := head, head
    for k > 1 {
        r = r.Next
        k--
    }
    for r != nil && r.Next != nil {
        l, r = l.Next, r.Next
    }
    return l
}
