/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    preH := &ListNode{Next: head}
    p1, p2, pre := head, head, preH
    // 先走n-1步
    for n > 0 && p1 != nil {
        p1 = p1.Next
        n--
    }

    // 找出目标节点
    for p1 != nil {
        p1, p2, pre = p1.Next, p2.Next, p2
    }

    pre.Next, p2.Next = p2.Next, nil


    return preH.Next
}
