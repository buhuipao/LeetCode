// 香农信息论
func poorPigs(buckets, minutesToDie, minutesToTest int) int {
    states := minutesToTest/minutesToDie + 1
    return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}
