func toLowerCase(s string) string {
    bs := []byte(s)
    for i := range bs {
        if bs[i] >= byte('A') && bs[i] <= byte('Z') {
            bs[i] += ('a' - 'A') 
        }
    }

    return string(bs)
}
