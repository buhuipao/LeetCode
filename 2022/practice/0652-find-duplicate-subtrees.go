/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    var ans []*TreeNode

    traverse(root, &ans, map[string]int{})

    return ans
}

func traverse(root *TreeNode, ret *[]*TreeNode, memo map[string]int) string {
    if root == nil {
        return "null"
    }

    left := traverse(root.Left, ret, memo)
    right := traverse(root.Right, ret, memo)
    s := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
    memo[s] += 1
    if memo[s] == 2 {
        *ret = append(*ret, root)
    }

    return s
}
