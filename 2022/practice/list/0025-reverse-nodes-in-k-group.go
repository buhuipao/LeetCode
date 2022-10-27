/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    var cnt int
    var p3, p4 *ListNode
    preHead := &ListNode{Next: head}
    p1, p2 := preHead, head
    for head != nil {
        if cnt == k {
            // 切断前后连接
            p1.Next, p3.Next = nil, nil
            // 反转
            reverse(p2)
            // 重新连接
            p1.Next, p2.Next = p3, p4

            p1, p2 = p2, p4
            cnt = 0
        }
        head, p3, p4 = head.Next, head, head.Next
        cnt++
   }
   if cnt == k {
        // 切断前后连接
        p1.Next, p3.Next = nil, nil
        // 反转
        reverse(p2)
        // 重新连接
        p1.Next, p2.Next = p3, p4
   }

    return preHead.Next
}

func reverse(head *ListNode) *ListNode {
    pre := new(ListNode)
    for head != nil {
        head, pre, head.Next = head.Next, head, pre
    }

    return pre
}e
    }

    return pre
}
