/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    nums := []int{}
    for i := 1; i <= n; i++ {
        nums = append(nums, i)
    }
    return helper(nums)
}

// 递归大法好
func helper(nums []int) []*TreeNode {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return []*TreeNode{&TreeNode{Val:nums[0]}}
    }
    ans := []*TreeNode{}
    for i := range nums {
        L, R := helper(nums[0:i]), helper(nums[i+1:])
        if len(L) == 0 {
            L = []*TreeNode{nil}
        }
        if len(R) == 0 {
            R = []*TreeNode{nil}
        }
        for j := range L {
            for k := range R {
                node := &TreeNode{Val: nums[i], Left: L[j], Right: R[k]}
                ans = append(ans, node)
            }
        }
    }
    return ans
}
