/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    h1, h2 := head, head.Next
    pre := &ListNode{Next: h1}
    H1, H2 := h1, h2
    // 确保下一轮的h1、h2都不为nil
    for h1 != nil && h1.Next != nil && h1.Next.Next != nil {
        h1.Next, h2.Next = h1.Next.Next, h2.Next.Next
        pre, h1, h2 = h1, h1.Next, h2.Next
    }
    // 如果h1不为空证明是奇数节点，直接用h1接H2，否则用pre接
    if h1 != nil {
        h1.Next = H2
    } else {
        pre.Next = H2
    }
    return H1
}
