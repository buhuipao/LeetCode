// 回溯
func countArrangement(n int) (ans int) {
    vis := make([]bool, n+1)
    match := make([][]int, n+1)
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i%j == 0 || j%i == 0 {
                match[i] = append(match[i], j)
            }
        }
    }

    var backtrack func(int)
    backtrack = func(index int) {
        if index > n {
            ans++
            return
        }
        for _, x := range match[index] {
            if !vis[x] {
                vis[x] = true
                backtrack(index + 1)
                vis[x] = false
            }
        }
    }
    backtrack(1)
    return
}
