# conding: utf-8
# 最长回文序列

'''
Given a string S, find the longest palindromic substring in S.
You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
'''


class Solution(object):
    def longestPalindrome(self, s):
        # 添加'#'，使的整个字符串变成奇数长度
        T = '#'.join(s)
        P = [0] * len(T)
        # 从中间查找，先查找左半部分，再右半部分，最后比较左右
        i = (len(T) + 1) // 2
        # 减少不必要验证取最大值即可
        while i >= max(P):
            while T[i + 1 + P[i]] == T[i - 1 - P[i]]:
                P[i] += 1
                # 防止左边越界
                if i - 1 - P[i] < 0:
                    break
            i -= 1
        i = (len(T) + 1) // 2 + 1
        while len(T) - i >= max(P):
            while T[i + 1 + P[i]] == T[i - 1 - P[i]]:
                P[i] += 1
                # 防止右边越界
                if i + 1 + P[i] == len(T):
                    break
            i += 1
        start = P.index(max(P)) - max(P)
        end = P.index(max(P)) + max(P)
        S = T[start: end + 1]
        return ''.join(S.split('#'))
