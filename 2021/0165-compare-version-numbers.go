// 双指针
func compareVersion(version1, version2 string) int {
    n, m := len(version1), len(version2)
    i, j := 0, 0
    for i < n || j < m {
        x := 0
        for ; i < n && version1[i] != '.'; i++ {
            x = x*10 + int(version1[i]-'0')
        }
        i++ // 跳过点号
        y := 0
        for ; j < m && version2[j] != '.'; j++ {
            y = y*10 + int(version2[j]-'0')
        }
        j++ // 跳过点号
        if x > y {
            return 1
        }
        if x < y {
            return -1
        }
    }
    return 0
}
