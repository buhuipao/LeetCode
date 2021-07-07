// 26进制
func convertToTitle(columnNumber int) string {
    ans := []byte{}

    for columnNumber > 0 {
        a0 := (columnNumber-1)%26 + 1
        ans = append(ans, 'A'+byte(a0-1))
        columnNumber = (columnNumber - a0) / 26
    }

    for i, n := 0, len(ans); i < n/2; i++ {
        ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
    }

    return string(ans)
}
