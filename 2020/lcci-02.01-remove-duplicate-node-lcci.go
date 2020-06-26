/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // O(1)空间，两层遍历
 // 另一种做法是hash
func removeDuplicateNodes(head *ListNode) *ListNode {
    var ans, node, pre *ListNode
    ans = head
    for head != nil {
        pre, node = head, head.Next
        for node != nil {
            // 摘除node，按住pre
            if node.Val == head.Val {
                pre.Next, node, node.Next = node.Next, node.Next, nil
            } else {
                // 遍历，替换pre、node
                pre, node = node, node.Next
            }
        }
        head = head.Next
    }
    return ans
}
