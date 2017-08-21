#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/08/20 14:13:35
#

'''
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
'''


class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        除了最后一个字母，其他都是左减右加
        """
        d = {'Ⅿ': 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
        result = 0
        for i in xrange(len(s)-1):
            if d[s[i]] > d[s[i+1]]:
                result += d[s[i]]
            else:
                result -= d[s[i]]
        return result + d[s[-1]]
