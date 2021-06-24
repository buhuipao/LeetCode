// 统计斜率数量
func maxPoints(points [][]int) (ans int) {
    n := len(points)
    if n <= 2 {
        return n
    }
    for i, p := range points {
        if ans >= n-i || ans > n/2 {
            break
        }
        cnt := map[int]int{}
        for _, q := range points[i+1:] {
            x, y := p[0]-q[0], p[1]-q[1]
            if x == 0 {
                y = 1
            } else if y == 0 {
                x = 1
            } else {
                if y < 0 {
                    x, y = -x, -y
                }
                g := gcd(abs(x), abs(y))
                x /= g
                y /= g
            }
            cnt[y+x*20001]++
        }
        for _, c := range cnt {
            ans = max(ans, c+1)
        }
    }
    return
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
