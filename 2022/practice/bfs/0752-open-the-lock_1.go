func openLock(deadends []string, target string) int {
    dead := make(map[int]bool)
    for i := range deadends {
        num, _ := strconv.Atoi(deadends[i])
        dead[num+10000] = true
    }

    const start = 10000
    targetN, _ := strconv.Atoi(target)
    targetN += 10000
    if targetN == start {
        return 0
    }
    if dead[start] {
        return -1
    }

    var step int
    q1 := map[int]bool{start: true}
    q2 := map[int]bool{targetN: true}
    seen := make(map[int]bool)
    for len(q1) > 0 && len(q2) > 0 {
        tmp := make(map[int]bool)
        if len(q2) < len(q1) {
            q1, q2 = q2, q1
        }

        for k := range q1 {
            if dead[k] {
                continue
            }

            if q2[k] { // 双向BFS有交集了，证明双向路径已通，可以返回
                return step
            }

            for _, nxt := range getNext(k) {
                if !seen[nxt] && !dead[nxt] {
                    if nxt == targetN {
                        return step + 1
                    }
                    tmp[nxt] = true
                }
            }
            seen[k] = true
        }
        step++
        q1, q2 = q2, tmp
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
