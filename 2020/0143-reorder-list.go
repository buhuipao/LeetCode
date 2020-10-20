/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    // 找到中点
    var pre *ListNode
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        pre, slow, fast = slow, slow.Next, fast.Next.Next
    }
    // 截断反转后半截
    pre.Next = nil 
    h1, h2 := head, reverse(slow)
    // 拼接两节，其中 h2 >= h1
    for h1 != nil {
        h1.Next, h1, h2, h2.Next, pre = h2, h1.Next, h2.Next, h1.Next, h2
    }
    if h2 != nil {
        pre.Next = h2
    }
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        head.Next, head, pre = pre, head.Next, head
    }
    return pre
}
