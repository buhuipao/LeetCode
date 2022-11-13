/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }
    
    preH := &ListNode{}
    l := preH
    // c 为进位
    var c, v, sum int
    for l1 != nil || l2 != nil || c != 0 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }

        sum = v1 + v2 + c
        c, v = sum / 10, sum % 10
        l.Next = &ListNode{Val: v}
        l = l.Next
    }

    return preH.Next
}
