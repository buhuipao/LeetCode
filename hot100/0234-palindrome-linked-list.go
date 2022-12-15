/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 双指针
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    pre := &ListNode{Next: head}
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        pre, slow, fast = pre.Next, slow.Next, fast.Next.Next
    }
    var head1, head2 *ListNode
    // 奇数节点
    if fast != nil {
        head1, pre.Next = pre, nil
        head2 = slow.Next
    // 偶数节点
    } else {
        head1, pre.Next = pre, nil
        head2 = slow
    }
    // 反转前半段link
    pre = new(ListNode)
    node := head
    for node != nil {
        node, pre, node.Next = node.Next, node, pre
    }
    for head1 != nil && head2 != nil {
        if head1.Val != head2.Val {
            return false
        }
        head1, head2 = head1.Next, head2.Next
    }
    return true
}
