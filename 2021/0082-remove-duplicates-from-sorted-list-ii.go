/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    preH := &ListNode{Val: 0, Next: head}
    cur := preH
    for cur.Next != nil && cur.Next.Next != nil {
        if cur.Next.Val == cur.Next.Next.Val {  // 发现了重复数值的节点，准备将之后的一段全部删除
            v := cur.Next.Val
            for cur.Next != nil && cur.Next.Val == v {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }
    return preH.Next
}
