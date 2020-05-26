/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    preH := &ListNode{Next: head}
    pre := preH
    for head != nil {
        if m <= 1 {
            break
        }
        head, pre = head.Next, head
        n--
        m--
    }
    node1 := pre
    pre.Next = nil
    oldH, H := head, head
    for ; head != nil; n-- {
        if n <= 1 {
            break
        }
        head, pre = head.Next, head
    }
    node2 := head.Next
    head.Next = nil
    node1.Next, oldH.Next = reverse(H), node2

    return preH.Next
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        pre, head, head.Next = head, head.Next, pre
    }
    return pre
}
