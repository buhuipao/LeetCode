type L struct {
    index int
    value int
}

func kWeakestRows(mat [][]int, k int) []int {
    lines := make([]L, 0, len(mat))
    for i := range mat {
        tmp := 0
        for j := range mat[i] {
            if mat[i][j] == 0 { // 军人都在前面
                break
            }
            tmp += mat[i][j]
        }
        lines = append(lines, L{i, tmp})
    }

    sort.Slice(lines, func(i, j int) bool {
        if lines[i].value == lines[j].value {
            return lines[i].index < lines[j].index
        }
        return lines[i].value < lines[j].value
    })

    fmt.Println(lines[:k])
    ans := make([]int, 0, k)
    for i := 0; i < k; i++ {
        ans = append(ans, lines[i].index)
    }
    
    return ans
}
