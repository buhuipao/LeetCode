#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/09/04 10:28:23
#

'''
Implement int sqrt(int x).

Compute and return the square root of x.
'''


class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        用牛顿迭代法
        f(r) = r ^ 2 - x, 求零点
        链接：https://www.zhihu.com/question/20690553
        """
        r = x
        while r ** 2 > x:
            r = (r + x / r) / 2
        return r
