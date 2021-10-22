func majorityElement(nums []int) (ans []int) {
    cnt := map[int]int{}
    for _, v := range nums {
        cnt[v]++
    }
    for v, c := range cnt {
        if c > len(nums)/3 {
            ans = append(ans, v)
        }
    }
    return
}
