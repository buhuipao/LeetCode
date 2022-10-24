func fib(n int) int {
    if n < 1 {
        return 0
    }

    dpA, dpB := 1, 1
    for i := 3; i <= n; i++ {
        dpA, dpB = dpB, (dpA + dpB) % 1000000007
    }
    
    return dpB
}
