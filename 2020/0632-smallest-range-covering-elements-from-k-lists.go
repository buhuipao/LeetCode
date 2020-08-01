// COPY FROM LEETCODE
// 自己想到的办法：1）把所有数组合并成一个长数组；2）按照滑动窗口的套路，并借助一个hash表不断寻常最优解
func smallestRange(nums [][]int) []int {
    size := len(nums)
    indices := map[int][]int{}
    xMin, xMax := math.MaxInt32, math.MinInt32
    for i := 0; i < size; i++ {
        for _, x := range nums[i] {
            indices[x] = append(indices[x], i)
            xMin = min(xMin, x)
            xMax = max(xMax, x)
        }
    }
    freq := make([]int, size)
    inside := 0
    left, right := xMin, xMin - 1
    bestLeft, bestRight := xMin, xMax
    for right < xMax {
        right++
        if len(indices[right]) > 0 {
            for _, x := range indices[right] {
                freq[x]++
                if freq[x] == 1 {
                    inside++
                }
            }
            for inside == size {
                if right - left < bestRight - bestLeft {
                    bestLeft, bestRight = left, right
                }
                for _, x := range indices[left] {
                    freq[x]--
                    if freq[x] == 0 {
                        inside--
                    }
                }
                left++
            }
        }
    }
    return []int{bestLeft, bestRight}
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

