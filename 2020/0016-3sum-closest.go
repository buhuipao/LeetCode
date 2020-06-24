import "sort"

func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)
    ans := nums[0] + nums[1] + nums[2]
    for i := range nums {
        // 跳过相同的数字
        for i > 0 && nums[i] == nums[i-1] {
            break
        }
        l, r := i+1, n-1
        for l < r {
            tmp := nums[i] + nums[l] + nums[r]
            // 符合要求直接返回
            if tmp == target {
                return target
            } else {
                v1, v2 := tmp - target, ans - target
                if v1 < 0 {
                    v1 *= -1
                }
                if v2 < 0 {
                    v2 *= -1
                }
                // 更新ans
                if v1 < v2 {
                    ans = tmp
                }
                if tmp > target {
                    // 跳过相同的l
                    for l < r && nums[r] == nums[r-1] {
                        r--
                    }
                    r--
                } else {
                    // 跳过相同的r
                    for l < r && nums[l] == nums[l+1] {
                        l++
                    }
                    l++
                }
            }
        }
    }
    return ans
}
