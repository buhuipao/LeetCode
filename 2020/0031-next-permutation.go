func nextPermutation(nums []int)  {
    n := len(nums)
    if n < 2 {
        return 
    }
    i := n - 2
    for ; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            break
        }
    }
    // 如果i合法，那么证明当前序列中存在递增序列，那么就需要先做一个替换
    // 比如：[1, 2, 3, 4, 2, 1]，因为从后往前看3、4是第一个递增序列，所以此时i为2，
    // 于是需要把3与其后面序列中第一个大于它的数字进行互换
    // 如果i不合法，那么就不需要以下步骤
    if i >= 0 {
        for j := n - 1; j > i; j-- {
            if nums[j] > nums[i] {
                nums[i], nums[j] = nums[j], nums[i]
                break
            }
        }
    }
    // 紧接上面的解释，兑换后num[i+1]（包含了i为-1的情况）～num[n-1]的序列仍然是单调逆序的，所以需要进行倒置
    l, r := i + 1, n - 1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
    return
}
