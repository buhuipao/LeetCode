func peakIndexInMountainArray(arr []int) int {
    ans, n := 1, len(arr)
    for i := 1; i < n-1; i++ {
        if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
            return i
        }
    }

    return ans
}
