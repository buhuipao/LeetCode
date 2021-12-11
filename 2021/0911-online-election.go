// 二分查找
type TopVotedCandidate struct {
    tops, times []int
}

func Constructor(persons, times []int) TopVotedCandidate {
    tops := []int{}
    top := -1
    voteCounts := map[int]int{-1: -1}
    for _, p := range persons {
        voteCounts[p]++
        if voteCounts[p] >= voteCounts[top] {
            top = p
        }
        tops = append(tops, top)
    }
    return TopVotedCandidate{tops, times}
}

func (c *TopVotedCandidate) Q(t int) int {
    return c.tops[sort.SearchInts(c.times, t+1)-1]
}
