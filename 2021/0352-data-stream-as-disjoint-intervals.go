 type SummaryRanges struct {
    *redblacktree.Tree
}

func Constructor() SummaryRanges {
    return SummaryRanges{redblacktree.NewWithIntComparator()}
}

func (ranges *SummaryRanges) AddNum(val int) {
    // 找到 l0 最大的且满足 l0 <= val 的区间 interval0 = [l0, r0]
    interval0, has0 := ranges.Floor(val)
    if has0 && val <= interval0.Value.(int) {
        // 情况一
        return
    }

    // 找到 l1 最小的且满足 l1 > val 的区间 interval1 = [l1, r1]
    // 在有序集合中，interval1 就是 interval0 的后一个区间
    interval1 := ranges.Iterator()
    if has0 {
        interval1 = ranges.IteratorAt(interval0)
    }
    has1 := interval1.Next()

    leftAside := has0 && interval0.Value.(int)+1 == val
    rightAside := has1 && interval1.Key().(int)-1 == val
    if leftAside && rightAside {
        // 情况四
        interval0.Value = interval1.Value().(int)
        ranges.Remove(interval1.Key())
    } else if leftAside {
        // 情况二
        interval0.Value = val
    } else if rightAside {
        // 情况三
        right := interval1.Value().(int)
        ranges.Remove(interval1.Key())
        ranges.Put(val, right)
    } else {
        // 情况五
        ranges.Put(val, val)
    }
}

func (ranges *SummaryRanges) GetIntervals() [][]int {
    ans := make([][]int, 0, ranges.Size())
    for it := ranges.Iterator(); it.Next(); {
        ans = append(ans, []int{it.Key().(int), it.Value().(int)})
    }
    return ans
}
