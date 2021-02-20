type entry struct{ cnt, l, r int }

// 字典记录
func findShortestSubArray(nums []int) (ans int) {
    mp := map[int]entry{}
    for i, v := range nums {
        if e, has := mp[v]; has {
            e.cnt++
            e.r = i
            mp[v] = e
        } else {
            mp[v] = entry{1, i, i}
        }
    }
    maxCnt := 0
    for _, e := range mp {
        if e.cnt > maxCnt {
            maxCnt, ans = e.cnt, e.r-e.l+1
        } else if e.cnt == maxCnt {
            ans = min(ans, e.r-e.l+1)
        }
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
