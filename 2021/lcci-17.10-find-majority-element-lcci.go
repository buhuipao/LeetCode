func majorityElement(nums []int) int {
    var ans, cnt int
    for _, num := range nums {
        if cnt == 0 {
            ans = num
        }
        if num == ans {
            cnt++
        } else {
            cnt--
        }
    }

    cnt = 0 
    for _, num := range nums {
        if num == ans {
            cnt++
        }
    }

    if 2 * cnt > len(nums) {
        return ans
    }

    return -1
}
