// Hash
func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize > 0 {
        return false
    }

    sort.Ints(hand)
    cnt := map[int]int{}
    for _, num := range hand {
        cnt[num]++
    }

    for _, x := range hand {
        if cnt[x] == 0 {
            continue
        }

        for num := x; num < x+groupSize; num++ {
            if cnt[num] == 0 {
                return false
            }
            cnt[num]--
        }
    }

    return true
}
