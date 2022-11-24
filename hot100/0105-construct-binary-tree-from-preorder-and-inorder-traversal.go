/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder, inorder []int, l1, r1, l2, r2 int) *TreeNode {
    if l1 > r1 {
        return nil
    }

    root := &TreeNode{Val: preorder[l1]}
    var index = l2
    for i := l2; i <= r2; i++ {
        if inorder[i] == preorder[l1] {
            index = i
            break
        }
    }

    root.Left = build(preorder, inorder, l1+1,  l1+(index-l2), l2, index-1)
    root.Right = build(preorder, inorder, l1+(index-l2)+1, r1, index+1, r2)

    return root
}
