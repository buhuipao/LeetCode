var options = map[string][]byte{
        "2": []byte{'a', 'b', 'c'},
        "3": []byte{'d', 'e', 'f'},
        "4": []byte{'g', 'h', 'i'},
        "5": []byte{'j', 'k', 'l'},
        "6": []byte{'m', 'n', 'o'},
        "7": []byte{'p', 'q', 'r', 's'}, 
        "8": []byte{'t', 'u', 'v'},
        "9": []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }

    var ans []string
    backtrack(digits, 0, []byte{}, &ans)

    return ans
}

func backtrack(s string, i int, path []byte, ans *[]string) {
    if i == len(s) {
        *ans = append(*ans, string(path))
        return
    }

    for _, v := range options[s[i:i+1]] {
        // 做出选择
        path = append(path, v)

        // 递归
        backtrack(s, i+1, path, ans)

        // 撤销选择
        path = path[:len(path)-1]
    }
}
