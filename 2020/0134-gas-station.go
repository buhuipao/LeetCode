func canCompleteCircuit(gas []int, cost []int) int {
    // 统计最终剩余量
    count := 0
    // 汽油缺口最大值以及其索引
    minV := int(^uint(0)>>1)
    minIndex := 0
    for i := range gas {
        count += gas[i] - cost[i]
        if count < minV {
            minV = count
            minIndex = i
        }
    }
    // 统计了所有加油站的油和整体消耗，发现最终剩余量小于零
    if count < 0 {
        return -1
    }
    // 如果总剩余大于0，那么就在最缺汽油的后一个点出发（也就是提前积累汽油，以便绕一周后能熬到最缺汽油的节点）
    return  (minIndex + 1) % len(gas)
}
