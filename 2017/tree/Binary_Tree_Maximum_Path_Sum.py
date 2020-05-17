# _*_ coding: utf-8 _*_

'''
Given a binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from \
some starting node to any node in the tree along the parent-child connections. \
The path must contain at least one node and does not need to go through the root.

For example:
Given the below binary tree,

       1
      / \
     2   3
Return 6.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        两个叶节点的和
        """
        self.max = float('-inf')
        self.pathHelper(root)
        return self.max

    def pathHelper(self, root):
        """
        返回root节点的和值更大的路径
        """
        if not root:
            return 0
        l = self.pathHelper(root.left)
        r = self.pathHelper(root.right)
        if l < 0: l = 0
        if r < 0: r = 0
        self.max = max(l + r + root.val, self.max)
        return max(root.val+l, root.val+r)
