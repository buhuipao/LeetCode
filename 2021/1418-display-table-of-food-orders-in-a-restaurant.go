func displayTable(orders [][]string) [][]string {
    foods := make(map[string]interface{}, len(orders)/8)
    tables := make(map[string]map[string]int, len(orders)/4)

    for _, o := range orders {
        _, table, food := o[0], o[1], o[2]
        foods[food] = nil
        if _, ok := tables[table]; !ok {
            tables[table] = make(map[string]int)
        }
        tables[table][food] += 1
    }

    // 收集并排序菜名
    headers := make([]string, 0, len(foods)+1)
    for f := range foods {
        headers = append(headers, f)
    }
    sort.Strings(headers)
    headers = append([]string{"Table"}, headers...)

    // 收集并排序座号
    sortedTables := make([]int, 0, len(tables))
    for t := range tables {
        tt, _ := strconv.Atoi(t)
        sortedTables = append(sortedTables, tt)
    }
    sort.Ints(sortedTables)

    ans := make([][]string, 0)
    ans = append(ans, headers)

    // 统计每桌点菜情况
    for _, tt := range sortedTables {
        t := strconv.Itoa(tt)
        v := tables[t]
        tmp := make([]string, 0, len(headers)+1)
        tmp = append(tmp, t)

        for _, f := range headers[1:] {
            tmp = append(tmp, strconv.Itoa( v[f]))
        } 

        ans = append(ans, tmp)
    }

    return ans
}
