func cuttingRope(n int) int {
    if n <= 3 {
        return n-1
    }

    a, b := n / 3, n % 3

    a--
    // https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/solution/jian-zhi-offer-14-ii-jian-sheng-zi-ii-by-xjwd/
    ans, v := 1, 3
    for a > 0 {
        if a % 2 == 1 {
            ans = (ans * v) % 1000000007
        }
        v = (v * v) % 1000000007
        a /= 2
    }

    if b == 1 {
        return (ans * 4) % 1000000007
    }

    if b == 2 {
        return (ans * 3 * 2) % 1000000007
    }

    return (ans * 3) % 1000000007
}
