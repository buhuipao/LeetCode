#!/usr/bin/env python
# _*_ coding: utf-8 _*_
#
#   Author:   buhuipao
#   E-mail:   chenhua22@outlook.com
#   Date:     17/09/03 18:17:42
#

'''
Given a non-negative integer represented as a non-empty array of digits,
plus one to the integer. You may assume the integer do not contain any leading
zero, except the number 0 itself. The digits are stored such that the most
significant digit is at the head of the list.
'''


class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        设置一个标志，如果标志在digits[0]处变了，证明需要进位需要digits需要插入0
        """
        leader = True
        for i in xrange(len(digits)-1, -1, -1):
            # 进位变0
            if digits[i] + 1 == 10:
                digits[i] = 0
            else:
                # 原位加1
                digits[i] += 1
                leader = False
                break
        if i == 0 and leader:
            digits.insert(0, 1)
        return digits
