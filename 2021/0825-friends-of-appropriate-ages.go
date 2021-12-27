// 排序双指针
func numFriendRequests(ages []int) (ans int) {
    sort.Ints(ages)
    left, right := 0, 0
    for _, age := range ages {
        if age < 15 {
            continue
        }
        for ages[left]*2 <= age+14 {
            left++
        }
        for right+1 < len(ages) && ages[right+1] <= age {
            right++
        }
        ans += right - left
    }
    return
}
