#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/08/19 18:45:24
#

'''
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

Note:
The input is assumed to be a 32-bit signed integer.
Your function should return 0 when the reversed integer overflows.
'''


class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        需要注意的是：数字的保存形式是补码，负数取余有点不同
        n位最大正数：0111...1 = 2 ** (n-1) - 1
        n位最小负数：100...0 = -2 ** (n-1), 而111...1 的值为-1
        """
        if x > 2 ** 31 - 1 or x < -(2 ** 31) + 1:
            return 0
        s = True
        if x < 0:
            x = -1 * x
            s = False
        o1 = []
        while x != 0:
            o1.append(x % 10)
            x = x / 10
        x = 0
        for o in o1:
            x = x * 10 + o
        if not s:
            x = -1 * x
        if x > 2 ** 31 - 1 or x < -(2 ** 31) + 1:
            return 0
        return x
