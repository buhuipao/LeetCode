/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    m, n := len(inorder), len(postorder)
    if m == 0 {
        return nil
    }
    root := &TreeNode{Val: postorder[n-1]}
    // 找到左子树的节点范围
    // inorder = [9,3,15,20,7]；postorder = [9,15,7,20,3]
    // 3为root，则左子树的范围为：inorder[0:1]或者postorder[0:1]
    // 右子树的范围为：inorder[2:]或者postorder[1:-1]
    var j int
    for i := range inorder {
        if inorder[i] == root.Val {
            j = i
            break
        }
    }
    root.Left = buildTree(inorder[0:j], postorder[0:j])
    root.Right = buildTree(inorder[j+1:], postorder[j:n-1])

    return root
}
