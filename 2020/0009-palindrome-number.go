func isPalindrome(x int) bool {
    // 特判
    if x < 0 || (x != 0 && x % 10 == 0){
        return false
    }
    num := 0
    for x > num {
         num = num * 10 + x % 10
         x = x / 10
    }
    // x为11时，1 == 1
    // x为121时，12 / 10 == 1
    return x == num || x == num / 10
}

/*
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
     nums := []int{}
    for x != 0 {
         nums = append(nums, x%10)
         x = x / 10
    }
    i, j := 0, len(nums)-1
    for i < j {
        if nums[i] != nums[j] {
            return false
        }
        i++
        j--
    }
    return true
}
*/
