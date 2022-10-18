func threeSum(nums []int) [][]int {
    return thressSumN(nums, 0)
}

func thressSumN(nums []int, target int) [][]int {
    sort.Ints(nums)
    var ans [][]int
    var tmp [][]int
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] { // 跳过相同项
            continue
        }

        tmp = twoSumN(nums, i+1, target-nums[i])
        for j := range tmp {
            ans = append(ans, []int{nums[i], tmp[j][0], tmp[j][1]})
        }
    }

    return ans
}

func twoSumN(nums []int, index, target int) [][]int {
    // 后续数字小于两个
    if index > len(nums)-2 {
        return nil
    }

    left, right := index, len(nums)-1
    var sum, leftV, rightV int
    var ans [][]int
    for left < right {
        leftV, rightV = nums[left], nums[right]
        sum = leftV + rightV
        if sum == target {
            ans = append(ans, []int{nums[left], nums[right]})
            for left < right && nums[left] == leftV {
                left++
            }
            for left < right && nums[right] == rightV {
                right--
            }
        } else if sum < target {
            for left < right && nums[left] == leftV {
                left++
            }
        } else {
            for left < right && nums[right] == rightV {
                right--
            }
        }
    }

    return ans
}
