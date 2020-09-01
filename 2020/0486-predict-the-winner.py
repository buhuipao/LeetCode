class Solution:
    def PredictTheWinner(self, nums: List[int]) -> bool:
        n = len(nums)
        dp = [[0]*n for _ in range(n)]

        for i in range(n):
            dp[i][i] = nums[i]

        for i in range(n-2, -1, -1):
            for j in range(i+1, n):
                dp[i][j] = max(nums[j]-dp[i][j-1], nums[i]-dp[i+1][j])

        return dp[0][-1]>=
