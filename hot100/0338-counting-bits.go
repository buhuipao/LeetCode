func countBits(num int) []int {
    ans := make([]int, num+1)
    for i := 1; i <= num; i++ {
        ans[i] = ans[i>>1] + i%2
    }
    
    return ans
}
