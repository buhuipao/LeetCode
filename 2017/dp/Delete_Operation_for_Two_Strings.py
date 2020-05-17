# _*_ coding: utf-8 _*_

'''
只要求出两个字符串的最长公共子序列，那么最终需要进行删除操作的就是m＋n－2*result
Given two words word1 and word2, find the minimum number of steps required
to make word1 and word2 the same, where in each step you can delete one character in either string.

Example 1:
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Note:
The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.
'''


class Solution(object):
    def minDistance1(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        LCS的变形题目
        """
        m, n = len(word1), len(word2)
        if not m or not n:
            return m or n
        dp = [[0] * (n+1) for i in xrange(m+1)]
        for i in xrange(m):
            for j in xrange(n):
                if word1[i] == word2[j]:
                    dp[i+1][j+1] = dp[i][j] + 1
                else:
                    dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
        return m + n - 2 * dp[i+1][j+1]
    
    def minDistance(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        节省了空间，时间复杂度没变
        """
        m, n = len(word1), len(word2)
        if not m or not n:
            return m or n
        dp1, dp2 = [0] * (n+1), [0] * (n+1)
        for i in xrange(m):
            dp1, dp2 = dp2, [0] * (n+1)
            for j in xrange(n):
                if word1[i] == word2[j]:
                    dp2[j+1] = dp1[j] + 1
                else:
                    dp2[j+1] = max(dp2[j], dp1[j+1])
        return m + n - 2 * dp2[j+1]
