func decode(encoded []int, first int) []int {
    ans := make([]int, len(encoded)+1)
    ans[0] = first
    for i := range encoded {
        ans[i+1] = encoded[i] ^ ans[i]
    }
    return ans
}
