func numTrees(n int) int {
    dp := make([]int, n+1)
    // dp[n]定义为n个有序数字可组成BST的数量
    // 空树也是一种BST
    dp[0], dp[1] = 1, 1
    for i := 2; i <= n; i++ {
        for j := 0; j < i; j++ {
            // 取[1..i]中一个数字为BST的root节点
            // 左子树分配j个节点，右子树分配i-j-1个节点
            dp[i] += dp[j] * dp[i-j-1]
        }
    }

    return dp[n]
}
