type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
    n := len(buildings)
    boundaries := make([]int, 0, n*2)
    for _, building := range buildings {
        boundaries = append(boundaries, building[0], building[1])
    }
    sort.Ints(boundaries)

    idx := 0
    h := hp{}
    for _, boundary := range boundaries {
        for idx < n && buildings[idx][0] <= boundary {
            heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
            idx++
        }
        for len(h) > 0 && h[0].right <= boundary {
            heap.Pop(&h)
        }

        maxn := 0
        if len(h) > 0 {
            maxn = h[0].height
        }
        if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
            ans = append(ans, []int{boundary, maxn})
        }
    }
    return
}
