func findRadius(houses, heaters []int) (ans int) {
    sort.Ints(houses)
    sort.Ints(heaters)
    j := 0
    for _, house := range houses {
        dis := abs(house - heaters[j])
        for j+1 < len(heaters) && abs(house-heaters[j]) >= abs(house-heaters[j+1]) {
            j++
            if abs(house-heaters[j]) < dis {
                dis = abs(house - heaters[j])
            }
        }
        if dis > ans {
            ans = dis
        }
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
