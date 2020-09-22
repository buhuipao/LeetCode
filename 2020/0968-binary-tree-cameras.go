/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 最小摄像头数量可以理解为物尽其用，从下层往上安装便可；
 // 因为从上往下安装，叶子数量过多，需要考虑覆盖的情况更多，很可能进行上层所采用的策略只是局部最优解；
 // 采用递归的思想
func minCameraCover(root *TreeNode) int {
    var ans int
    // 0：此节点未覆盖
    // 1：此节点已被其他节点摄像头覆盖
    // 2：此节点已直接安装
    var h func(*TreeNode) int
    h = func (root *TreeNode) int {
        if root == nil {
            return 1
        }
        L, R := h(root.Left), h(root.Right)
        // 当前节点需要安装一个摄像头对某一个子节点进行覆盖，以便节省一个摄像头
        if L == 0 || R == 0 {
            ans++
            return 2
        }
        // 两个子节点都被覆盖了
        if L == 1 && R == 1 {
            return 0
        }
        // L+R >= 3
        // 由于子节点原因，自身已经被覆盖了
        return 1
    }

    if h(root) == 0 {
        ans++
    }

    return ans
}
