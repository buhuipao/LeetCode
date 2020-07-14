// 递归 + 原地修改
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    // 没有值
    if n == 0 {
        return 0
    }
    // 只剩下最后一层，返回顶点
    if n == 1 {
        return triangle[0][0]
    }
    // 将最后一行的值计算进倒数第二层里去
    for i := range triangle[n-2] {
        triangle[n-2][i] += min(triangle[n-1][i], triangle[n-1][i+1])
    }
    
    return minimumTotal(triangle[:n-1])
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

/* 
还可以O(N) DP解决，因为每个节点的最小和只取决于上层的两个值，存上一层的memo行了
*/
