func sumNums(n int) int {
    ans := 0
    var sum func(int) bool
    sum = func(n int) bool {
        ans += n
        return n > 0 && sum(n-1)    // 递归终止条件
    }
    sum(n)
    return ans
}
