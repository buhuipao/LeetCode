func openLock(deadends []string, target string) int {
    dead := make(map[int]bool)
    for i := range deadends {
        num, _ := strconv.Atoi(deadends[i])
        dead[num+10000] = true
    }

    const start = 10000
    targetN, _ := strconv.Atoi(target)
    if targetN+10000 == start {
        return 0
    }
    if dead[start] {
        return -1
    }

    type pair struct {
        status int 
        step   int
    }
    q := []pair{{start, 0}}
    seen := map[int]bool{start: true}
    for len(q) > 0 {
        p := q[0]
        q = q[1:]
        for _, nxt := range getNext(p.status) {
            if !seen[nxt] && !dead[nxt] {
                if nxt == targetN + 10000 {
                    return p.step + 1
                }
                seen[nxt] = true
                q = append(q, pair{nxt, p.step + 1})
            }
        }
    }
    return -1
}

// 枚举 status 通过一次旋转得到的数字
func getNext(status int) []int {
    var nums []int
    var tmp int
    for i := 0; i < 4; i++ {
        tmp, status = status % 10, status / 10
        nums = append(nums, tmp)
    }

    var ret []int
    var n, v int
    for i := 0; i < 4; i++ {
        n = nums[i]
        // 向左
        if n == 0 {
            nums[i] = 9
        } else {
            nums[i] = n - 1
        }
        v = 0
        for j := 3; j >= 0; j-- {
            v = v * 10 + nums[j] 
        }
        nums[i] = n
        ret = append(ret, v+10000)

        // 向右
        if n == 9 {
            nums[i] = 0
        } else {
            nums[i] = n + 1
        }
        v = 0
        for j := 3; j >= 0; j-- {
            v = v * 10 + nums[j] 
        }
        nums[i] = n
        ret = append(ret, v+10000)
    }

    return ret
}
