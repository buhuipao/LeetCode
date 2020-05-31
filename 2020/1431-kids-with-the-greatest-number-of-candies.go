func kidsWithCandies(candies []int, extraCandies int) []bool {
    n := len(candies)
    if n == 0 {
        return nil
    }
    ans := make([]bool, n)
    max := candies[0]
    for i := 0; i < n; i++ {
        if candies[i] > max {
            max = candies[i]
        }
    }
    for i := 0; i < n; i++ {
        if candies[i] + extraCandies >= max {
            ans[i] = true
        }
    }
    return ans
}
