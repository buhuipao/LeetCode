func networkDelayTime(times [][]int, n, k int) (ans int) {
    const inf = math.MaxInt64 / 2
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = inf
        }
    }
    
    for _, t := range times {
        x, y := t[0]-1, t[1]-1
        g[x][y] = t[2]
    }

    dist := make([]int, n)
    for i := range dist {
        dist[i] = inf
    }
    dist[k-1] = 0
    used := make([]bool, n)
    for i := 0; i < n; i++ {
        x := -1
        for y, u := range used {
            if !u && (x == -1 || dist[y] < dist[x]) {
                x = y
            }
        }
        used[x] = true
        for y, time := range g[x] {
            dist[y] = min(dist[y], dist[x]+time)
        }
    }

    for _, d := range dist {
        if d == inf {
            return -1
        }
        ans = max(ans, d)
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
