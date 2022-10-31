/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(inorder, postorder []int, l1, r1, l2, r2 int) *TreeNode {
    if l1 > r1 {
        return nil
    }

    root := &TreeNode{Val: postorder[r2]}
    var index = r1
    for i := r1; i >= l1; i-- {
        if inorder[i] == root.Val {
            index = i
            break
        }
    }

    root.Left = build(inorder, postorder, l1, index-1, l2, l2+(index-l1)-1)
    root.Right = build(inorder, postorder, index+1, r1, r2-(r1-index),r2-1)

    return root
}

