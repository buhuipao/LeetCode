func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	// 标记合数
	arr := make([]bool, n+1)
	for i := 2; i*i < n; i++ {
		if arr[i] {
			continue
		}

		for j := i * i; j < n; j += i {
			arr[j] = true
		}
	}

	var ans int
	for i := 2; i < n; i++ {
		if !arr[i] {
			ans++
		}
	}

	return ans
}