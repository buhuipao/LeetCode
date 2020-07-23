/* https://leetcode-cn.com/problems/divisor-game/solution/python3gui-na-fa-by-pandawakaka/
func divisorGame(N int) bool {
    return N % 2 == 0
}
*/

// DP做法
func divisorGame(N int) bool {
    if N < 2 {
        return false
    }
    dp := make([]bool, N+1)
    dp[2] = true
    for i := 3; i <= N; i++ {
        for j := 1; j < i/2; j++ {
            // 证明存在一种可能，使得爱丽丝选j时，满足条件且对方输
            if i % j == 0 && !dp[i-j] {
                dp[i] = true
                break
            }
        }
    }
    return dp[N]
}
