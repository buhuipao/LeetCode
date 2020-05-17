#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/09/17 00:06:12
#

"""
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3,
the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].

Note:
Try to come up as many solutions as you can,
there are at least 3 different ways to solve this problem.
"""

class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        if nums and k:
            n = len(nums)
            k = n - k % n
            for i in xrange(k):
                nums.append(nums.pop(0))

    def rotate1(self, nums, k):
        n = len(nums)
        k = k % n
        if k:
            nums[:k], nums[k:] = nums[-k:], nums[:-k]
