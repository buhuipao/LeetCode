func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)
    for _, s := range ransomNote {
        m[s] += 1
    }

    for _, s := range magazine {
        if _, ok := m[s]; ok {
            m[s] -= 1
        }
        if m[s] == 0 {
            delete(m, s)
        }
    }

    return len(m) == 0 
}
