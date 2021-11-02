/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    var pre *ListNode
    for node.Next != nil {
        node.Val, node, pre = node.Next.Val, node.Next, node
    }

    pre.Next = nil
}
