class Solution:
    def new21Game(self, N: int, K: int, W: int) -> float:
        """
        dp[i]为游戏过程中可能到达点数i的概率，应是之前W个概率之和除以W，需要注意边界条件。
        """
        dp = [1] * (N + 1)
        cur = 0
        for i in range(1, N + 1):
            if i <= K:
                cur += dp[i - 1]
            if K + W >= i > W:
                cur -= dp[i - 1 - W]
            print(dp)
            dp[i] = cur / W
        return sum(dp[K:])
