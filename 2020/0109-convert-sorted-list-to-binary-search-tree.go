/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    a, b, c := cutList(head)
    L, R := sortedListToBST(a), sortedListToBST(c)
    return &TreeNode{Val: b.Val, Left: L, Right: R}
}

func cutList(head *ListNode) (*ListNode, *ListNode, *ListNode) {
    // 只有一个节点
    if head.Next == nil {
        return nil, head, nil
    }
    var pre, fast, slow, next *ListNode
    fast, slow = head, head
    for fast != nil && fast.Next != nil {
        pre, fast, slow = slow, fast.Next.Next, slow.Next
    }
    next = slow.Next
    // 切断
    pre.Next, slow.Next = nil, nil
    return head, slow, next
}
