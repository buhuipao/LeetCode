// 回溯
func generateParenthesis(n int) []string {
    if n == 0 {
        return nil 
    }

    var ans []string
    backtrack(nil, n, n, &ans)
    return ans
}

func backtrack(path []byte, l, r int, ans *[]string) {
    if l == 0 && r == 0 {
        *ans = append(*ans, string(path))
        return
    }

    // 右括号已用数量大于左括号，例如：())，则明显不符合要求
    if r < l {
        return
    }

    if l > 0 {
        path = append(path, '(')
        backtrack(path, l-1, r, ans)
        path = path[:len(path)-1]
    }

    if r > 0 {
        path = append(path, ')')
        backtrack(path, l, r-1, ans)
        path = path[:len(path)-1]
    }
}
