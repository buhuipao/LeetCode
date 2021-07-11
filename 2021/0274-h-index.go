// 排序
func hIndex(citations []int) int {
    sort.Ints(citations)

    var h int
    for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
        h++
    }

    return h
}
