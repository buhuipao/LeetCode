# Link: https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        preResult ListNode
        up int
        value1, value2 int
    )
    p := &preResult
    for l1 != nil || l2 != nil || up != 0 {
        value1, value2 = 0, 0
        if l1 != nil {
            value1 = l1.Val
        }
        if l2 != nil {
            value2 = l2.Val
        }
        value := (value1 + value2 + up)
        up = value / 10
        p.Next = &ListNode{Val: value % 10}
        p = p.Next
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }

    return preResult.Next
}
