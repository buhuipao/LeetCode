/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    preH := &ListNode{
        Next: head,
    }
    var pre, target *ListNode
    // 先把target放在-1的位置，然后让head先跑n-1步
    target = preH
    for head != nil {
        if n > 1 {
            n--
        } else {
            target, pre = target.Next, target
        }
        head = head.Next
    }
    // 删除target
    pre.Next = target.Next

    return preH.Next
}
