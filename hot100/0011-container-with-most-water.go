// 双指针
func maxArea(height []int) int {
    l, r := 0, len(height)-1
    var h, ans int
    for l < r {
        width := r - l
        // 每次都移动短板
        if height[l] > height[r] {
            h = height[r]
            r--
        } else {
            h = height[l]
            l++
        }

        if width * h > ans {
            ans = width * h
        }
    }

    return ans
}
