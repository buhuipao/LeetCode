# _*_ coding: utf-8 _*_

'''
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along \
the shortest path from the root node down to the nearest leaf node.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def minDepth(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        l = self.minDepth(root.left)
        r = self.minDepth(root.right)
        # 对于非叶子节的单子树节点，必须取更大的子树高度而不是空子树的高度0
        # 其他的情况就直接取高度更少的树
        return (min(l, r) or max(l, r)) + 1
