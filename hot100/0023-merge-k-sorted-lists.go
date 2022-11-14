/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    if len(lists) == 2 {
        return mergeTwoList(lists[0], lists[1])
    }

    return mergeTwoList(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    preH := &ListNode{}
    pre := preH
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            l1, pre.Next = l1.Next, l1
        } else {
            l2, pre.Next = l2.Next, l2
        }

        pre = pre.Next
    }

    if l1 != nil {
        pre.Next = l1
    }

    if l2 != nil {
        pre.Next = l2
    }

    return preH.Next
}
