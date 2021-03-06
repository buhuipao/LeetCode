const (
    P = math.MaxInt32
    BASE = 41
)
// DP：dp[i] = min(dp[i], dp[i-j])
func respace(dictionary []string, sentence string) int {
    hashValues := map[int]bool{}
    for _, word := range dictionary {
        hashValues[getHash(word)] = true
    }
    // 初始化，设定最坏值
    f := make([]int, len(sentence) + 1)
    for i := 1; i < len(f); i++ {
        f[i] = len(sentence)
    }
    for i := 1; i <= len(sentence); i++ {
        f[i] = f[i-1] + 1
        hashValue := 0
        for j := i; j >= 1; j-- {
            t := int(sentence[j-1] - 'a') + 1
            hashValue = (hashValue * BASE + t) % P
            if hashValues[hashValue] {
                f[i] = min(f[i], f[j-1])
            }
        }
    }
    return f[len(sentence)]
}

// 螺旋hash：https://zh.wikipedia.org/wiki/Rabin%E2%80%93Karp%E7%AE%97%E6%B3%95
// 利用前一个位置的hash值计算出当前位置的hash值
func getHash(s string) int {
    hashValue := 0
    for i := len(s) - 1; i >= 0; i-- {
        hashValue = (hashValue * BASE + int(s[i] - 'a') + 1) % P
    }
    return hashValue
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

