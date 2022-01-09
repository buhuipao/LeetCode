func slowestKey(releaseTimes []int, keysPressed string) byte {
    ans := keysPressed[0]
    maxTime := releaseTimes[0]
    for i := 1; i < len(keysPressed); i++ {
        key := keysPressed[i]
        time := releaseTimes[i] - releaseTimes[i-1]
        if time > maxTime || time == maxTime && key > ans {
            ans = key
            maxTime = time
        }
    }
    return ans
}
