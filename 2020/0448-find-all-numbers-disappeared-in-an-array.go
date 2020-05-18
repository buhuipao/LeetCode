// 利用了鸽子洞原理
// 把nums中可以读取到的数字（小于或者等于n）作为索引号，
// 尝试把所有可以通过这些索引号访问到的数字进行标记,
// 如何标记？1）原数字*-1，这样换算为索引时乘-1即可；2）原数字+n，同时换算索引时取模
// 最后没有被标记的位置就是缺少的数字
func findDisappearedNumbers(nums []int) []int {
    for i := range nums {
        index := nums[i]
        if index < 0 {
            index *= -1
        }
        if nums[index-1] > 0 {
            nums[index-1] *= -1
        }
    }
    ans := []int{}
    for i := range nums {
        if nums[i] > 0 {
            ans = append(ans, i+1)
        }
    }
    return ans
}
