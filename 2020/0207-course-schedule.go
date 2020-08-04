// 拓扑排序&是否有环
func canFinish(numCourses int, prerequisites [][]int) bool {
    // 入度表
    in := make(map[int][]int, numCourses)
    table := make([]int, numCourses)
    for i := range prerequisites {
        des, src := prerequisites[i][0], prerequisites[i][1]
        in[src] = append(in[src], des)
        table[des]++
    }
    // 找到初始入度为0的点
    queue := make([]int, 0, numCourses)
    for i := range table {
        if table[i] == 0 {
            queue = append(queue, i)
        }
    }
    var count int
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        count++
        for i := range in[node] {
            table[in[node][i]]--
            if table[in[node][i]] == 0 {
                queue = append(queue, in[node][i])
            }
        }
    }
    return count == numCourses
}
