func threeSum(nums []int) [][]int {
    return threeSumN(nums, 0)
}

func threeSumN(nums []int, target int) [][]int {
    sort.Ints(nums)
    var ans [][]int
    var tmp [][]int
    for i := 0; i < len(nums); i++ {
        tmp = twoSumN(nums, i+1, target-nums[i])
        for j := range tmp {
            ans = append(ans, append([]int{nums[i]}, tmp[j]...))
        }

        for i < len(nums)-1 && nums[i] == nums[i+1] {
            i++
        }
    }

    return ans
}

func twoSumN(nums []int, index, target int) [][]int {
    // 后续剩余数字小于2
    if index > len(nums)-2 {
        return nil
    }
    
    l, r := index, len(nums)-1
    var ans [][]int
    for l < r {
        lV, rV := nums[l], nums[r]
        if lV + rV == target {
            ans = append(ans, []int{lV, rV})
            for l < r && nums[l] == lV {
                l++
            }
            for l < r && nums[r] == rV {
                r--
            }
        } else if lV + rV < target {
            for l < r && nums[l] == lV {
                l++
            }
        } else {
            for l < r && nums[r] == rV {
                r--
            }
        }
    }

    return ans
}
