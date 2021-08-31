// 暴力法
// 还有差分法，类比于：统计离散时间点里一辆公交车上的人数
func corpFlightBookings(bookings [][]int, n int) []int {
    ans := make([]int, n)
    for _, b := range bookings {
        i, j, v := b[0], b[1], b[2]
        for i <= j {
            ans[i-1] += v
            i++
        }
    }

    return ans
}
