/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    pre := &ListNode{Next: head}
    preH := pre
    for head != nil {
        tail := head
        for i := 1; i < k && tail != nil; i++ {
            tail = tail.Next
        }
        if tail == nil {
            break
        }
        // 找到下一个头，并截断当前的尾巴
        nextHead := tail.Next
        tail.Next = nil
        // 连接新的头
        pre.Next, head.Next = reverse(head), nextHead
        pre, head = head, nextHead
    }
    return preH.Next
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        pre, head, head.Next = head, head.Next, pre
    }
    return pre
}
