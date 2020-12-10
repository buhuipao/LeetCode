func lemonadeChange(bills []int) bool {
    m := make(map[int]int, 0)
    for _, b := range bills {
        if b == 10 {
            if m[5] == 0 {
                return false
            }
            m[5] -= 1
        } else if b == 20 {
            if m[5] == 0 {
                return false
            }
            if m[10] > 0 {
                m[10] -= 1
                m[5] -= 1
            } else if m[5] >= 3{
                m[5] -= 3
            } else {
                return false
            }
        }
        m[b] += 1
    }
    return true
}
