/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    // 统计长度
    var n int
    node := head
    for node != nil {
        n++
        node = node.Next
    }
    // 切换并拼接
    var l1, l2 *ListNode
    l2 = head
    for m := n - (k%n); m > 1; m-- {
        l2 = l2.Next
    }
    l1, l2.Next = l2.Next, nil  // 切断
    if l1 != nil {
        node = l1
        for node.Next != nil {
            node = node.Next
        }
        node.Next = head
        return l1
    }
    return head
}
