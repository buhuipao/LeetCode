# _*_ coding: utf-8 _*_


class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        # 先进行排序
        nums, result, i = sorted(nums), [], 0
        while i < len(nums) - 2:
            # 除了第一项，之后谨防出现相同后项
            if i == 0 or nums[i] != nums[i - 1]:
                # 固定一位，往中间夹
                j, k = i + 1, len(nums) - 1
                while j < k:
                    if nums[i] + nums[j] + nums[k] < 0:
                        j += 1
                    elif nums[i] + nums[j] + nums[k] > 0:
                        k -= 1
                    else:
                        result.append([nums[i], nums[j], nums[k]])
                        j, k = j + 1, k - 1
                        # 当初出现符合条件的项时，及时过滤之后的想同项目
                        while j < k and nums[j] == nums[j - 1]:
                            j += 1
                        while j < k and nums[k] == nums[k + 1]:
                            k -= 1
            i += 1
        return result
