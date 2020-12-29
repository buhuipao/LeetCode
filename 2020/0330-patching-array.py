class Solution:
    def minPatches(self, nums: List[int], n: int) -> int:
        furthest = i = ans = 0
        while furthest < n:
            # 可覆盖到，直接用前缀和更新区间
            if i < len(nums) and nums[i] <= furthest + 1:
                furthest += nums[i] #  [1, furthest] -> [1, furthest + nums[i]]
                i += 1
            else:
                # 不可覆盖到，增加一个数 furthest + 1，并用前缀和更新区间
                furthest = 2 * furthest + 1 # [1, furthest] -> [1, furthest + furthest + 1]
                ans += 1
        return ans
