// 差分数组
func corpFlightBookings(bookings [][]int, n int) []int {
    ans := make([]int, n)
    for i := range bookings {
        ans[bookings[i][0]-1] += bookings[i][2]
        if bookings[i][1] < n {
            ans[bookings[i][1]] -= bookings[i][2]
        }
    }

    for i := 1; i < n; i++ {
        ans[i] += ans[i-1]
    }

    return ans
}
