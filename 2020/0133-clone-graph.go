/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
// DFS
func cloneGraph(node *Node) *Node {
    visited := make(map[int]*Node)
    return dfs(visited, node)
}

func dfs(visited map[int]*Node, node *Node) *Node {
    // 空节点
    if node == nil {
        return nil
    }
    // 已经访问过
    if n, ok := visited[node.Val]; ok {
        return n
    }
    newN := &Node{Val: node.Val, Neighbors: []*Node{}}
    visited[newN.Val] = newN
    for _, n := range node.Neighbors {
        newN.Neighbors = append(newN.Neighbors, dfs(visited, n))
    }
    return newN
}
