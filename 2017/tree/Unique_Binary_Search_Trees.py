# _*_ coding: utf-8 _*_

'''
Given n, how many structurally unique BST's (binary search trees) that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
'''


class Solution(object):
    def numTrees(self, n):
        """
        :type n: int
        :rtype: int
        每次以i为root节点，那么求n就是以1, 2, 3...n为root节点的总和
        但必须先构建n之前节点的树，求出1～(n-1)的结果
        """
        if n < 0:
            return 0
        dp = [0] * (n + 1)
        # 空树和一个节点
        dp[0] = dp[1] = 1
        i = 2
        while i <= n:
            # 构建n之前的树
            for j in xrange(1, i+1):
                dp[i] += dp[j-1] * dp[i-j]
            i += 1
        return dp[n]
