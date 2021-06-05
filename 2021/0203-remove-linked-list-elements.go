/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    preH := &ListNode{Next: head}
    pre := preH
    for head != nil {
        if head.Val == val {
            pre.Next = head.Next
        } else {
            pre = head
        }
        head = head.Next
    }
    return preH.Next
}
