/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) > 2 {
        l1, l2 := mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:])
        return mergeKLists([]*ListNode{l1, l2})
    }

    if len(lists) == 0 {
        return nil
    }

    if len(lists) == 1 {
        return lists[0]
    }

    // 合并两条
    return mergeLists(lists[0], lists[1])
}

func mergeLists(l1, l2 *ListNode) *ListNode {
    preRoot := &ListNode{}
    root := preRoot

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            root.Next, l1 = l1, l1.Next
        } else {
            root.Next, l2 = l2, l2.Next
        }

        root = root.Next
    }

    if l1 != nil {
        root.Next = l1
    } else {
        root.Next = l2
    }

    return preRoot.Next
}
