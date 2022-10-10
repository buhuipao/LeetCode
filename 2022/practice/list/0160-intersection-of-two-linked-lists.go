/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 1）逻辑相交：将两个链表逻辑层面拼接遍历；
 // 2）计数：统计长度后遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    p1, p2 := headA, headB
    for p1 != p2 {
        if p1 == nil {
            p1 = headB
        } else {
            p1 = p1.Next
        }

        if p2 == nil {
            p2 = headA
        } else {
            p2 = p2.Next
        }
    }

    return p1
}
