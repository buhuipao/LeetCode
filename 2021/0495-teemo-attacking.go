func findPoisonedDuration(timeSeries []int, duration int) int {
    var sum, end int
    n := len(timeSeries)
    for i := 0; i < n-1; i++ {
        end = timeSeries[i] + duration - 1
        sum += min(0, timeSeries[i+1] - end - 1) + duration
    }

    sum += duration

    return sum
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
