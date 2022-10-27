/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    pre, oldHead := &ListNode{Next: head}, head
    for head != nil {
        pre, head, head.Next = head, head.Next, pre 
    }
    oldHead.Next = nil

    return pre
}
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    // 只有一个节点
    if head.Next == nil {
        return head
    }

    newHead := reverseList(head.Next)
    // 反转head和head.Next，其实也就是连接反转完成链表的尾巴和当前节点
    newTail := head.Next
    newTail.Next = head
    // 将最后的尾巴置为空
    head.Next = nil

    return newHead
}
