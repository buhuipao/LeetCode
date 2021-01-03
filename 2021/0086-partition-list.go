/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    h1, h2 := &ListNode{}, &ListNode{}
    preH1, preH2 := h1, h2
    for head != nil {
        if head.Val < x {
            h1.Next = head
            h1 = h1.Next
        } else {
            h2.Next = head
            h2 = h2.Next
        }
        head, head.Next = head.Next, nil
    }
    h1.Next = preH2.Next
    return preH1.Next
}
