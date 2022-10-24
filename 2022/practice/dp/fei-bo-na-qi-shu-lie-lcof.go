func fib(n int) int {
    if n < 1 {
        return 0
    }

    memo := make(map[int]int, n)
    return helper(memo, n)
}

func helper(memo map[int]int, n int) int {
    if n == 1 || n == 2 {
        return 1
    }

    if memo[n] != 0 {
        return memo[n]
    }

    memo[n] = (helper(memo, n-1) + helper(memo, n-2)) % 1000000007

    return memo[n]
}
