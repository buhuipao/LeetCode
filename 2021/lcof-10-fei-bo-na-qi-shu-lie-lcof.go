func fib(n int) int {
    a, b := 0, 1
    for i := 0; i < n; i++ {
        a, b = b, (a + b) % 1000000007
    }
    return a
}
