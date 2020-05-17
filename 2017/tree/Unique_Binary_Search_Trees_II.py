# _*_ coding: utf-8 _*_

'''
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1...n.

For example,
Given n = 3, your program should return all 5 unique BST's shown below.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def generateTrees(self, n):
        """
        :type n: int
        :rtype: List[TreeNode]
        基本思路就是：
        每次以i点为root节点，1至i－1为左子树，i+1至n为右子树, 同理在左右子树中也可以如此递归
        每个以i为root的树的种类为：G(i) ＝G(i-1) * G(n-i)
        """
        def generate(m, n):
            # 考虑到最左最右的情况
            if m > n:
                return [None]
            if m == n:
                return [TreeNode(m)]

            result = []
            for i in xrange(m, n+1):
                for l in generate(m, i-1):
                    for r in generate(i+1, n):
                        root = TreeNode(i)
                        root.left, root.right = l, r
                        result.append(root)
            return result
        if n == 0:
            return []
        return generate(1, n)
