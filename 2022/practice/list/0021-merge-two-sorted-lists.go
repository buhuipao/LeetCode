/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    root := &ListNode{}
    preRoot := root
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            root.Next, list1 = list1, list1.Next
        } else {
            root.Next, list2 = list2, list2.Next
        }

        root = root.Next
    }
    
    if list1 != nil {
        root.Next = list1
    } else {
        root.Next = list2
    }

    return preRoot.Next
}
