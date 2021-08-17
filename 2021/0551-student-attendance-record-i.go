func checkRecord(s string) bool {
    var c1, c2 int
    for i := range s {
        if s[i] == 'A' {
            c1++
        }
        if c1 == 2 {
            return false
        }

        if s[i] == 'L' {
            c2++
        } else {
            c2 = 0
        }
        if c2 == 3 {
            return false
        }

    }
    
    return true
}
