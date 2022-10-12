/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 快慢指针
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    slow, fast := head, head
    for fast != nil {
        if slow.Val != fast.Val {
            slow.Next, slow = fast, fast
        }

        fast = fast.Next
    }

    // slow.Next = nil 用于切断后续重复节点
    slow.Next = nil

    return head
}
