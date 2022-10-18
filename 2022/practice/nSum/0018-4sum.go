func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    return nSum(nums, 4, 0, target)
}

func nSum(nums []int, n, index, target int) [][]int {
    length := len(nums)
    // 递归的尽头
    if n < 2 {
        return nil
    }

    // 剩余元素不够了，直接返回
    if length - index < n {
        return nil
    }

    var ret, tmp [][]int
    if n > 2 {
        for ; index < length; index++ {
            tmp = nSum(nums, n-1, index+1, target-nums[index])
            for i := range tmp {
                ret = append(ret, append([]int{nums[index]}, tmp[i]...))
            }
            // 跳过重复项
            for index < length-1 && nums[index] == nums[index+1] {
                index++
            }
        }
    } else { // two-sum
        var left, right, sum int
        l, r := index, length-1
        for l < r {
            left, right, sum = nums[l], nums[r], nums[l]+nums[r]
            if sum < target {
                for l < r && nums[l] == left {
                    l++
                }
            } else if sum > target {
                for l < r && nums[r] == right {
                    r--
                }
            } else {
                ret = append(ret, []int{left, right})
                for l < r && nums[l] == left {
                    l++
                }
                for l < r && nums[r] == right {
                    r--
                }
            }
        }
    }

    return ret
}
