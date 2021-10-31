
const (
    A = `qwertyuiop`
    B = `asdfghjkl`
    C = `zxcvbnm`
)


func findWords(words []string) []string {
    m := make(map[byte]int, 26*2)
    v := byte('a' - 'A')

    for i := range A {
        m[A[i]] = 1
        m[A[i]-v] = 1
    }

    for i := range B {
        m[B[i]] = 2 
        m[B[i]-v] = 2
    }

    for i := range C {
        m[C[i]] = 3
        m[C[i]-v] = 3
    }

    ans := make([]string, 0)
    for _, w := range words {
        flag := 0
        for i := range w {
            if flag == 0 {
                flag = m[w[i]]
                continue
            }

            if flag != m[w[i]] {
                flag = 0
                break
            }
        }

        if flag != 0 {
            ans = append(ans, w)
        }
    }

    return ans
}
