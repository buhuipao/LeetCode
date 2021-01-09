# 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/yi-kan-jiu-dong-de-jian-dan-ti-jie-by-ed-0o3l/
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0
        # 5种状态
        dp0 = 0             # 一直不买
        dp1 = - prices[0]   # 到最后也只买入了一笔
        dp2 = float("-inf") # 到最后买入一笔，卖出一笔
        dp3 = float("-inf") # 到最后买入两笔，卖出一笔
        dp4 = float("-inf") # 到最后买入两笔，卖出两笔

        for i in range(1, len(prices)):
            dp1 = max(dp1, dp0 - prices[i])
            dp2 = max(dp2, dp1 + prices[i])
            dp3 = max(dp3, dp2 - prices[i])
            dp4 = max(dp4, dp3 + prices[i])

        return max(dp0, dp2, dp4)
