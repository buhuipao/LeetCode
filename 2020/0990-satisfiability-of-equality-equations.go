// 并查集
func equationsPossible(equations []string) bool {
    // 初始化各个独立的节点
    parent := make([]int, 26)
    for i := 0; i < 26; i++ {
        parent[i] = i
    }
    // 根据相等的情况union各个节点
    for _, str := range equations {
        if str[1] == '=' {
            index1 := int(str[0] - 'a')
            index2 := int(str[3] - 'a')
            union(parent, index1, index2)
        }
    }
    // 检查矛盾
    for _, str := range equations {
        if str[1] == '!' {
            index1 := int(str[0] - 'a')
            index2 := int(str[3] - 'a')
            // 如果两个应该不相等的节点之前是有相等关联，那么证明矛盾了
            if find(parent, index1) == find(parent, index2) {
                return false
            }
        }
    }
    return true
}

func union(parent []int, index1, index2 int) {
    parent[find(parent, index1)] = find(parent, index2)
}

func find(parent []int, index int) int {
    // 压缩路径，返回root
    for parent[index] != index {
        parent[index] = parent[parent[index]]
        index = parent[index]
    }
    return index
}
