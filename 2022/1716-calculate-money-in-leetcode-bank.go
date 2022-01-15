// 两次等差数列
func totalMoney(n int) (ans int) {
    // 所有完整的周存的钱
    weekNum := n / 7
    firstWeekMoney := (1 + 7) * 7 / 2
    lastWeekMoney := firstWeekMoney + 7*(weekNum-1)
    weekMoney := (firstWeekMoney + lastWeekMoney) * weekNum / 2
    
    // 剩下的不能构成一个完整的周的天数里存的钱
    dayNum := n % 7
    firstDayMoney := 1 + weekNum
    lastDayMoney := firstDayMoney + dayNum - 1
    dayMoney := (firstDayMoney + lastDayMoney) * dayNum / 2
    return weekMoney + dayMoney
}
