// 回溯
func wordBreak(s string, wordDict []string) []string {
    var ans []string
    backtrack(s, 0, wordDict, nil, &ans)
    return ans
}

func backtrack(s string, i int, wordDict []string, path []string, ans *[]string) {
    // 终止条件
    if i == len(s) {
        tmp := make([]string, len(path))
        copy(tmp, path)
        *ans = append(*ans, strings.Join(tmp, " "))
        return
    }

    for _, w := range wordDict {
        // 剪枝
        if i + len(w) > len(s) {
            continue
        }
        if w != s[i:i+len(w)] {
            continue
        }

        path = append(path, w)
        backtrack(s, i+len(w), wordDict, path, ans)
        path = path[:len(path)-1]
    }
}
