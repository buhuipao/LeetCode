// 排序，两头取值求和
func numRescueBoats(people []int, limit int) (ans int) {
    sort.Ints(people)
    light, heavy := 0, len(people)-1
    for light <= heavy {
        if people[light]+people[heavy] > limit {
            heavy--
        } else {
            light++
            heavy--
        }
        ans++
    }
    return
}
