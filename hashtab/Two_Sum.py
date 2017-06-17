# _*_ coding: utf-8 _*_

'''
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
'''


class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        假设有唯一解，可以用hash表存下target-num的索引
        time: O(n)
        space: O(n)
        """
        idxDict = dict()
        for idx, num in enumerate(nums):
            if target - num in idxDict:
                return [idxDict[target - num], idx]
            idxDict[num] = idx

    def twoSum2(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        先排序再用夹逼法进行查找
        """
        _nums = []
        for i in xrange(len(nums)):
            _nums.append((i, nums[i]))
        _nums = sorted(_nums, key=lambda num: num[1])
        i, j = 0, len(_nums)-1
        while i < j:
            if _nums[i][1] + _nums[j][1] < target:
                i += 1
            elif _nums[i][1] + _nums[j][1] > target:
                j -= 1
            else:
                return [_nums[i][0], _nums[j][0]]
