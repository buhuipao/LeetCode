/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
    if l > r {
        return nil
    }

    index := -1
    maxValue := ^int(^uint(0)>>1)
    for i := l; i <= r; i++ {
        if nums[i] > maxValue {
            index, maxValue = i, nums[i] 
        }
    }

    root := &TreeNode{Val: maxValue}
    root.Left = build(nums, l, index-1)
    root.Right = build(nums, index+1, r)

    return root
}
