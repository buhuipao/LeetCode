/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    } 

    var D int
    for _, c := range root.Children {
        if tmp := maxDepth(c); tmp > D {
            D = tmp
        }
    }

    return D + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
