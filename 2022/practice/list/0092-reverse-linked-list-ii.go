/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return nil
    }

    preHead := &ListNode{Next: head}
    pre := preHead
    var cnt int
    for left > cnt+1 && head != nil {
        head, pre = head.Next, head
        cnt++
    }
    newTail, newHeadPre := head, pre
    pre.Next = nil

    for cnt+1 < right && head != nil {
        head, pre = head.Next, head
        cnt++
    }
    newTailNext := head.Next
    head.Next = nil

    newHead := reverse(newTail)
    newHeadPre.Next, newTail.Next = newHead, newTailNext


    return preHead.Next
}

func reverse(head *ListNode) *ListNode {
    pre := new(ListNode)
    for head != nil {
        pre, head, head.Next = head, head.Next, pre
    }

    return pre
}
