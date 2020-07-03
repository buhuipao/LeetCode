/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
     n := len(nums)
    // 退出条件
    if n == 0 {
        return nil
    }
    if n == 1 {
        return &TreeNode{Val: nums[0]}
    }
    L, R := sortedArrayToBST(nums[:n/2]), sortedArrayToBST(nums[n/2+1:])
    return &TreeNode{Val: nums[n/2], Left: L, Right: R}
}
