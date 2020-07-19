func twoSum(numbers []int, target int) []int {
    L, R := 0, len(numbers)-1
    var ans []int
    for L < R {
       tmp := numbers[L] + numbers[R]
       if tmp < target {
           L++
       } else if tmp == target {
           ans = []int{L+1, R+1}
           break
       } else {
           R--
       }
    }
    return ans
}
