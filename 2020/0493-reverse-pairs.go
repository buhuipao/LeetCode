func reversePairs(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }
    // 切分拷贝数组
    n1 := append([]int{}, nums[:n/2]...)
    n2 := append([]int{}, nums[n/2:]...)
    // 递归完毕后，n1 和 n2 均为有序
    // 1、分别统计n1、n2内的翻转对
    cnt := reversePairs(n1) + reversePairs(n2)
    // 2、统计跨n1、n2的翻转对
    // 这里比较巧妙，由于 n1 和 n2 均为有序，借用了n1[i+1] > n1[i]直接继承上一个的j
    j := 0
    for _, v := range n1 {
        for j < len(n2) && v > 2*n2[j] {
            j++
        }
        cnt += j
    }
    // n1 和 n2 归并填入 nums
    p1, p2 := 0, 0
    for i := range nums {
        // 如果n1还没遍厉完同时（n2遍历完了或者n1中的数字更小）
        if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
            nums[i] = n1[p1]
            p1++
        } else {
            nums[i] = n2[p2]
            p2++
        }
    }
    return cnt
}
