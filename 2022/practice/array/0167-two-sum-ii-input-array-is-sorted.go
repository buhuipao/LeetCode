// 双指针
func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers)-1
    var tmp int
    for l < r {
        tmp = numbers[l] + numbers[r]
        if tmp == target {
            return []int{l+1, r+1}
        }

        if tmp < target {
            l++
        } else {
            r--
        }
    }

    return nil
}
