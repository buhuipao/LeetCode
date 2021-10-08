func findRepeatedDnaSequences(s string) (ans []string) {
    L := 10
    cnt := map[string]int{}
    for i := 0; i <= len(s)-L; i++ {
        sub := s[i : i+L]
        cnt[sub]++
        if cnt[sub] == 2 {
            ans = append(ans, sub)
        }
    }
    
    return
}
