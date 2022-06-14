/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}
*/

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var stack1, stack2 []*TreeNode
    stack1 = append(stack1, root)
    for len(stack1) != 0 {
        node := stack1[len(stack1)-1]
        stack1 = stack1[:len(stack1)-1]
        stack2 = append(stack2, node)

        if node.Left != nil {
            stack1 = append(stack1, node.Left)
        }

        if node.Right != nil {
            stack1 = append(stack1, node.Right)
        }
    }

    ans := make([]int, 0, len(stack2))
    for len(stack2) != 0 {
        ans = append(ans, stack2[len(stack2)-1].Val)
        stack2 = stack2[:len(stack2)-1] 
    }

    return ans
}
