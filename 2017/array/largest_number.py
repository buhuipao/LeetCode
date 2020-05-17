# _*_ coding: utf-8 _*_

'''
Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

Note: The result may be very large, so you need to return a string instead of an integer.
'''


class Solution:
    # @param {integer[]} nums
    # @return {string}
    def largestNumber(self, nums):
        nums = sorted(map(str, nums))
        # 自定义比较函数，
        # cmp(y+x, x+y)如果返回－1，证明y+x更小(y<x)，然后x(较大)就排在y之前
        nums.sort(cmp=lambda x, y: cmp(y + x, x + y))
        # 假如左边全为'0', 就只返回一个'0'
        result = ''.join(nums).lstrip('0') or '0'
        return result
