/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 递归·分治·快慢指针
func sortList(head *ListNode) *ListNode {
    // 临界条件
    // 最多只有一个节点的链表
    if head == nil || head.Next == nil {
        return head
    }
    // 找到中间节点
    slow, fast := head, head
    var preS *ListNode
    for fast != nil && fast.Next != nil {
        preS = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 根据中间节点把链表一切为二分别递归排序，最终合并即可
    preS.Next = nil
    left := sortList(head)
    right := sortList(slow)
    
    return helper(left, right)
}

// 合并两个排序好的链表
func helper(l1, l2 *ListNode) *ListNode {
    // 假设l1、l2都是排序好的
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    pre := &ListNode{}
    node := pre
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            node.Next = l2
            l2 = l2.Next
        } else {
            node.Next = l1
            l1 = l1.Next
        }
        node = node.Next
    }
    if l1 != nil {
        node.Next = l1
    }
    if l2 != nil {
        node.Next = l2
    }
    return pre.Next
}
