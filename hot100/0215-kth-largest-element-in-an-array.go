// 快速选择
func findKthLargest(nums []int, k int) int {
    l, r := 0, len(nums)-1
    // 转换为索引
    index := len(nums) - k
    var p int
    for l <= r {
        p = partition(nums, l, r)
        if p < index {
            l = p + 1
        } else if p > index {
            r = p - 1
        } else {
            return nums[p]
        }
    }

    return -1
}

func partition(nums []int, l, r int) int {
    if l == r {
        return l
    }

    // 选择l为默认分界点
    pivot := nums[l]
    // 这里i为l+1-1，是因为在其for循环中会先加一
    // 这里j为r+1，是因为在其for循环中会先减一
    // !!!为什么要在进循环前减一呢？ 因为会出现nums[i或者j] == pivot的情况，为了避免死循环
    i, j := (l+1)-1, r+1
    for {
        // nums[l+1:i]都是小于pivot
        for i++; nums[i] < pivot; i++ {
            if i == r {
                break
            }
        }
        // nums[j+1:r+1]都大于pivot
        for j--; nums[j] > pivot; j-- {
            if j == l {
                break
            }
        }

        // 已经越界
        if i >= j {
            break
        }

        // 交换i和j
        nums[i], nums[j] = nums[j], nums[i]
    }

    // 交换pivot到正确位置
    nums[j], nums[l] = nums[l], nums[j]

    // 现在nums[l:j]... < nums[j] < nums[j+1:r+1]...
    // 返回pivot处于的位置j
    return j
}
