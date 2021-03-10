func subarraysWithKDistinct(A []int, K int) (ans int) {
    n := len(A)
    num1 := make([]int, n+1)
    num2 := make([]int, n+1)
    var tot1, tot2, left1, left2 int
    for _, v := range A {
        if num1[v] == 0 {
            tot1++
        }
        num1[v]++
        if num2[v] == 0 {
            tot2++
        }
        num2[v]++
        for tot1 > K {
            num1[A[left1]]--
            if num1[A[left1]] == 0 {
                tot1--
            }
            left1++
        }
        for tot2 > K-1 {
            num2[A[left2]]--
            if num2[A[left2]] == 0 {
                tot2--
            }
            left2++
        }
        ans += left2 - left1
    }
    return ans
}
