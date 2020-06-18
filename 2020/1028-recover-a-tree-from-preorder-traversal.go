func recoverFromPreorder(S string) *TreeNode {
    path, pos := []*TreeNode{}, 0
    for pos < len(S) {
        level := 0
        for S[pos] == '-' {
            level++
            pos++
        }
        value := 0
        for ; pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos++ {
            value = value * 10 + int(S[pos] - '0')
        }
        node := &TreeNode{Val: value}
        // 如果level和stack内的节点数相同，证明是上一个节点的左节点
        if level == len(path) {
            // root节点特殊处理
            if len(path) > 0 {
                path[len(path)-1].Left = node
            }
        } else {
            // 回溯到node的父节点
            path = path[:level] 
            path[len(path)-1].Right = node
        }
        // 把所有节点都放入stack中
        path = append(path, node)
    }
    return path[0]
}
