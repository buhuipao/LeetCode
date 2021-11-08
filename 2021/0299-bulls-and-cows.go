// 遍历
func getHint(secret, guess string) string {
    bulls := 0
    var cntS, cntG [10]int
    for i := range secret {
        if secret[i] == guess[i] {
            bulls++
        } else {
            cntS[secret[i]-'0']++
            cntG[guess[i]-'0']++
        }
    }
    cows := 0
    for i := 0; i < 10; i++ {
        cows += min(cntS[i], cntG[i])
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
