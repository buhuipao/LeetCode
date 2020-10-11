class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        # 基础检查
        if not nums:
            return False
        n = len(nums)
        if n < 2:
            return False
        
        total = sum(nums)
        target = total //2
        if total % 2 != 0:
            return False

        # 记忆化递归
        dp = [set() for _ in range(n)]
        def recursion(nums, i, remain):
            if remain == 0:
                return True
            # 剩余值小于0、数组遍历完了、已经计算过了并且还没退出返回则同样不要计算了
            if remain < 0 or i == n or remain in dp[i]:
                return False
            dp[i].add(remain)
            #  选择当前nums[i]或者不选
            return recursion(nums, i+1, remain- nums[i]) or recursion(nums, i+1, remain)

        return recursion(nums, 0, target)
