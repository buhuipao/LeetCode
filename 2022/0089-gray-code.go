// 1）镜像法：{0|X}、{1|X}
// 2）公式： G(i) = i ^ (i/2); 
// Python一行：[i ^ i >> 1  for i in range(2 ** n)]
func grayCode(n int) []int {
    ans, head := make([]int, 0, 1 << n), 1 // head等于1意味着已经左移一次了，可以直接进入循环了
    ans = append(ans, 0)
    for i := 1; i <= n; i++ {
        for j := len(ans)-1; j >= 0; j-- {
            ans = append(ans, ans[j] + head)
        }
        head = head << 1
    }
    return ans
}
