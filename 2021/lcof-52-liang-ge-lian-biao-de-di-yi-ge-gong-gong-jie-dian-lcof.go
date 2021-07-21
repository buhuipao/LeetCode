/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var a, b int
    A, B := headA, headB
    for A != nil || B != nil {
        if A != nil {
            A = A.Next
            a++
        }
        if B != nil {
            B = B.Next
            b++
        }
    }
    // 假设A比B长
    A, B = headA, headB
    if a < b {
        A, B = B, A
        a, b = b, a
    }
    for a > b && A != nil {
        A = A.Next
        a--
    }
    for A != nil {
        if A == B {
            return A
        }
        A, B = A.Next, B.Next
    }
    return nil
}
