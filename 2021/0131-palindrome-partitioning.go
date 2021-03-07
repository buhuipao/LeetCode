// 回溯 + 减枝
func partition(s string) [][]string {
    if len(s) == 0 {
        return nil
    }
    // memo记录已经验证过的字符串是否为回文
    // 也有些做法是建一个二维dp，dp[i][j]表示s[i:j+1]是否为回文，
    // dp的状态转移方程为：dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
    // 而我这里采用了hash表的方式记录会带来两个问题：
    // 1）更多的函数调用，整体变慢；2）字符串copy传递和hash的记录，导致内存占用会变多
    memo := make(map[string]bool, 0)
    ans := make([][]string, 0)
    path := make([]string, 0)
    backtrack(path, &s, 0, &ans, memo)
    return ans
}

func backtrack(path []string, options *string, index int, ans *[][]string, memo map[string]bool) {
    n := len(*options)
    // 结束条件：已经到字符串末尾了
    if index == n {
        *ans = append(*ans, path)
        return
    }
    for i := index; i < n; i++ {
        // 减枝
        if !isPalindrome(options, index, i+1, memo) {
            continue
        }
        // 做出选择
        newP := make([]string, len(path))
        copy(newP, path)
        newP = append(newP, (*options)[index:i+1])
        // 进入下一步决策
        backtrack(newP, options, i+1, ans, memo)
        // 撤销选择
        //
    }
}

// 判断字符串是否为回文串，同时记录到备忘录里
func isPalindrome(s *string, start, end int, memo map[string]bool) bool {
    if _, ok := memo[(*s)[start:end]]; ok {
        return memo[(*s)[start:end]]
    }
    i, j := start, end - 1
    for i <= j {
        if (*s)[i] != (*s)[j] {
            memo[(*s)[start:end]] = false
            return false
        }
        i++
        j--
    }
    memo[(*s)[start:end]] = true
    return true
}
