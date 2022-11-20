func trap(height []int) int {
    l, r := 0, len(height)-1
    var maxL, maxR, ans int
    for l < r {
        // 每次都找到左右最高的墙
        maxL = max(height[l], maxL)
        maxR = max(height[r], maxR)
        // 基于左右墙中较矮的那一个往前推进
        // 因为每一列积水的高度取决于左右墙中较矮的那一面，所以需要积极找出左右墙中较矮的那一个
        if maxL < maxR {
            ans += maxL - height[l]
            l++
        } else {
            ans += maxR - height[r]
            r--
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
