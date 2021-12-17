func numWaterBottles(numBottles int, numExchange int) int {
    bottle, ans := numBottles, numBottles
    for bottle >= numExchange {
        bottle = bottle - numExchange
        ans++
        bottle++
    }
    return ans
}
