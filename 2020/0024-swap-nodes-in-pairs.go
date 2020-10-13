/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    pre := &ListNode{Next: head}
    preH := pre
    for head != nil && head.Next != nil {
        nextH := head.Next.Next

        pre.Next, head.Next.Next, head.Next = head.Next, head, nextH

        head, pre = nextH, head
    }

    return preH.Next
}
