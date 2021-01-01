func canPlaceFlowers(flowerbed []int, n int) bool {
    // 为了处理首位两个0的特殊情况，但这么处理会导致内存占用变大一点
    flowerbed, cnt := append(flowerbed, 0), 1
    N := len(flowerbed)
    // 查找连续3个以上的零
    for i := 0; i < N && n > 0; i++ {
        if flowerbed[i] == 0 {
            cnt++
        } else {
            cnt = 0
        }
        // 超过3个零就可以种花，且种完花重置为1
        if cnt == 3 {
            n--
            cnt = 1
        }
    }
    return n == 0
}
