import (
    "sort"
)

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for i := range strs {
        t := []byte(strs[i])
        sort.Slice(t, func(i, j int) bool {
            return t[i] < t[j]
        })
        s := string(t)

        if _, ok := m[s]; !ok {
            m[s] = make([]string, 0, 1)
        }
        
        m[s] = append(m[s], strs[i])
    }

    ans := make([][]string, 0, len(m))
    for _, v := range m {
        ans = append(ans, v)
    }

    return ans
}
