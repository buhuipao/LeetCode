/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 中序遍历
func findMode(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    stack := []*TreeNode{root}
    node := root.Left
    cur, max := -1, 0
    pre := root.Val + 1
    ans := make([]int, 0)
    for node != nil || len(stack) != 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }

        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 重置cur值，更新max值
        if node.Val != pre {
            // 清空历史ans，更新最大值
            if cur > max {
                ans = []int{pre}
                max = cur
            } else if cur == max {
                ans = append(ans, pre)
            } // 忽略小于max的
            cur = 1
        } else {
            cur++
        }

        node, pre = node.Right, node.Val
    }
    // 最后一个点
    if cur > max {
        ans = []int{pre}
    } else if cur == max {
        ans = append(ans, pre)
    }

    return ans
}
