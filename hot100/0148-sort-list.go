/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    fast, slow := head, head
    var pre *ListNode
    for fast != nil && fast.Next != nil {
        fast, slow, pre = fast.Next.Next, slow.Next, slow
    }
    // 切断pre和slow
    pre.Next = nil

    return merge(sortList(head), sortList(slow))
}

func merge(l1, l2 *ListNode) *ListNode {
    preH := &ListNode{}
    pre := preH
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }

    if l1 != nil {
        pre.Next = l1
    }

    if l2 != nil {
        pre.Next = l2
    }

    return preH.Next
}
