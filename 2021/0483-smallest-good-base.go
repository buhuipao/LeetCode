func smallestGoodBase(n string) string {
    nVal, _ := strconv.Atoi(n)
    // 利用数学求和公式算出最大值
    mMax := bits.Len(uint(nVal)) - 1
    for m := mMax; m > 1; m-- {
        k := int(math.Pow(float64(nVal), 1/float64(m)))
        mul, sum := 1, 1
        for i := 0; i < m; i++ {
            mul *= k
            sum += mul
        }
        if sum == nVal {    // 尝试并得到答案
            return strconv.Itoa(k)
        }
    }
    return strconv.Itoa(nVal - 1)
}
