func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    for i := range strs {
        cnt := [26]int{}
        for j := range strs[i] {
            cnt[strs[i][j]-'a']++
        }
        m[cnt] = append(m[cnt], strs[i])
    }

    ans := make([][]string, 0, len(m))
    for _, v := range m {
        ans = append(ans, v)
    }

    return ans
}
