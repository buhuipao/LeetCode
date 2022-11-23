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

// 递归
func helper(nums []int) []*TreeNode {
   if len(nums) == 0 {
       return []*TreeNode{nil}
   } 

   if len(nums) == 1 {
       return []*TreeNode{{Val: nums[0]}}
   }
   
    var ans []*TreeNode
    for i := range nums {
        L, R := helper(nums[:i]), helper(nums[i+1:])
        for j := range L {
            for k := range R {
                ans = append(ans, &TreeNode{Val: nums[i], Left: L[j], Right: R[k]})
            }
        }
    } 

    return ans
}
