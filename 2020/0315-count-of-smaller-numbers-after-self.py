class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        ans = []
        sorted_nums = []
        # 从右往左遍历，并建立有序数组，利用二分查找找到每个数的索引值
        for i in range(len(nums) - 1, -1, -1):
            idx = bisect.bisect_left(sorted_nums, nums[i])  # 二分查找，Python自带的函数
            ans.append(idx)
            sorted_nums.insert(idx, nums[i])
        return ans[::-1]
