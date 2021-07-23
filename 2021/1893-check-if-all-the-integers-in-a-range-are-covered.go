func isCovered(ranges [][]int, left, right int) bool {
    diff := [52]int{} // 差分数组
    for _, r := range ranges {
        diff[r[0]]++
        diff[r[1]+1]--
    }
  
    cnt := 0 // 前缀和
    for i := 1; i <= 50; i++ {
        cnt += diff[i]
        if cnt <= 0 && left <= i && i <= right {
            return false
        }
    }
  
    return true
}
