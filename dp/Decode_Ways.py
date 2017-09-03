#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/08/20 22:01:37
#

'''
A message containing letters from A-Z is being encoded to numbers using
the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given an encoded message containing digits,
determine the total number of ways to decode it.

For example,
Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

The number of ways decoding "12" is 2.

'''


class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        (d>'0')*w: 保证了当前的数字可以独立成为一个字母
        (9<int(p+d)<27)*v: 保证了前一个和当前可以组合被译码为一个字母
        """
        v, w, p = 0, int(s > ''), ''
        for d in s:
            v, w, p = w, (d > '0') * w + (9 < int(p+d) < 27) * v, d
        return w
