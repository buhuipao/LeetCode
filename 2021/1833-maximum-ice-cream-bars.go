func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)

    var ans int
    for i := range costs {
        if coins < costs[i] {
            return ans
        }
        coins -= costs[i]
        ans++
    }

    return ans
}
