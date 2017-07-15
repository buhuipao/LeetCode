# _*_ coding: utf-8 _*_

'''
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from \
the root node down to the farthest leaf node.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def maxDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        r = self.maxDepth(root.right)
        l = self.maxDepth(root.left)
        return max(r, l) + 1
