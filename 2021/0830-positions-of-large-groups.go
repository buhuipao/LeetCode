func largeGroupPositions(s string) [][]int {
    n, cnt := len(s), 1
    ans := make([][]int, 0)
    for i := 1; i < n; i++ {
        if s[i] == s[i-1] {
            cnt++
        } else {
            if cnt >= 3 {
                ans = append(ans, []int{i-cnt,i-1})
            }
            cnt = 1
        }
    } 
    // 补充最后一个重复子字符串
    if cnt >= 3 {
        ans = append(ans, []int{n-cnt, n-1})
    }
    return ans
}
