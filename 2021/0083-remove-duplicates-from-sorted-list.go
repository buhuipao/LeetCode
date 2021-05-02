/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    node := head
    pre := &ListNode{Next: node}
    preH := pre
    for node != nil {
        for node != nil && node.Next != nil && node.Val == node.Next.Val { // 只保留最后一个节点
            node = node.Next
        }
        pre.Next = node
        pre, node = node, node.Next
    }
    return preH.Next
}
