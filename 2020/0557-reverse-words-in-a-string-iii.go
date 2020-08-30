import (
    "bytes"
    "strings"
)

func reverseWords(s string) string {
    buffer := bytes.Buffer{}
    buffer.Grow(len(s))
    strs := strings.Split(s, " ")
    for i := range strs {
        n := len(strs[i])
        for j := n-1; j >= 0; j-- {
            buffer.WriteByte(strs[i][j])
        }
        buffer.WriteByte(' ')
    }
    // 截断最后一个多余的空格
    buffer.Truncate(len(s))
    
    return buffer.String()
}
