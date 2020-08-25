// 回溯
// 有点类似全排列，应该是多全排列的拼接版
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    ans := make([]string, 0)
    options := map[string][]string{
        "2": []string{"a", "b", "c"},
        "3": []string{"d", "e", "f"},
        "4": []string{"g", "h", "i"},
        "5": []string{"j", "k", "l"},
        "6": []string{"m", "n", "o"},
        "7": []string{"p", "q", "r", "s"}, 
        "8": []string{"t", "u", "v"},
        "9": []string{"w", "x", "y", "z"},
    }
    backtrack(digits, 0, "", options, &ans)
    return ans
}

func backtrack(s string, i int, path string, options map[string][]string, ans *[]string) {
    // 终止条件
    if i == len(s) {
        *ans = append(*ans, path)
        return
    }
    for j := range options[s[i:i+1]] {
        // 作出选择
        var buffer bytes.Buffer
        buffer.WriteString(path)
        buffer.WriteString(options[s[i:i+1]][j])
        // 递归
        backtrack(s, i+1, buffer.String(), options, ans)
        // 撤销选择
        //
    }
}
