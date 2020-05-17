// 拓扑排序
func findOrder(numCourses int, prerequisites [][]int) []int {
    // 特判
    if numCourses == 0 {
        return nil
    }
    // 建立出度表
    outTable := make(map[int][]int, numCourses)
    inCount := make(map[int]int, numCourses)
    for i := range prerequisites {
        des, src := prerequisites[i][0], prerequisites[i][1]
        if _, ok := outTable[src]; !ok {
            outTable[src] = make([]int, 0)
        }
        inCount[des] += 1
        outTable[src] = append(outTable[src], des)
    }
    // BFS队列，初始化入度为0的队列
    queue := make([]int, 0)
    for i := 0; i < numCourses; i++ {
        if inCount[i] == 0 {
            queue = append(queue, i)
        }
    }
    ans := []int{}
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        ans = append(ans, node)
        // 处理
        for _, v := range outTable[node] {
            // 可以加入入度为零的队列了
            if inCount[v] == 1 {
                queue = append(queue, v)
            }
            inCount[v] -= 1
        }
    }
    // 有环，提前出了循环
    if len(ans) != numCourses {
        return nil
    }
    return ans
}
