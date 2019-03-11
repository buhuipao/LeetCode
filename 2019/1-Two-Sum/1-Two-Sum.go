// Link: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
    // 使用map
    tmpMap, result := make(map[int]int), make([]int, 2)
    for i := range nums {
        if tmp, ok := tmpMap[target - nums[i]]; ok {
            result = []int{i, tmp}
            break
        }
        tmpMap[nums[i]] = i
    }
    return result
}
