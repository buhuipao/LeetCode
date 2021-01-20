class Solution:
    def maximumProduct(self, nums: List[int]) -> int:
        nums = sorted(nums)
        a, b, c = nums[-1], nums[-2], nums[-3]
        x, y = nums[0], nums[1]
        return max(a * b * c, x * y * a)
