func minArray(numbers []int) int {
    left, right := 0, len(numbers)-1
    for left < right {
        mid := left + (right - left) / 2
        // 证明最小值在(mid, right]，如下：
        // [2, 3, 4, 1, 1]，mid为2，rigth为4
        if numbers[mid] > numbers[right] {
            left = mid + 1
        // 证明最小值在[left, mid]，如下：
        // [6, 2, 3, 4, 5]，mid为2，right为4
        } else if numbers[mid] < numbers[right] {
            right = mid
        // 无法确定最小值在哪个区间，如下三种情况：
        // 1）[1, 2, 2, 3]，mid为1，right为2
        // 2）[2, 2, 2, 1, 2, 2]，mid为2， rigth为5
        // 3）[2, 2, 2, 2, 2, 2]
        // 但是可以肯定的是，最小值一定是在right的左边
        } else {
            right--
        }
    }
    return numbers[left]
}
