func fourSum(nums []int, target int) (quadruplets [][]int) {
    // 排序
    sort.Ints(nums)
    n := len(nums)
    for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
        // 跳过重复 或者 右边不符合条件的
        if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
            continue
        }
        for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
            // 跳过重复的 或者 右边不符合条件的
            if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
                continue
            }
            // 两数之和
            for left, right := j+1, n-1; left < right; {
                if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
                    quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
                    // 跳过重复
                    for left++; left < right && nums[left] == nums[left-1]; left++ {
                    }
                    for right--; left < right && nums[right] == nums[right+1]; right-- {
                    }
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return
}
