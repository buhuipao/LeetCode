# _*_ coding: utf-8 _*_

'''
Given an array with n objects colored red, white or blue,
sort them so that objects of the same color are adjacent,
with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red,
white, and blue respectively.

Note:
You are not suppose to use the library's sort function for this problem.
'''


class Solution(object):
    def sortColors(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        # 计数排序
        if not nums:
            return nums
        r = nums.count(0)
        w = nums.count(1)
        b = nums.count(2)
        count = {'0': r, '1': w, '2': b}

        n = 0
        for i in xrange(3):
            nums[n:n+count[str(i)]] = [i] * count[str(i)]
            n += count[str(i)]
