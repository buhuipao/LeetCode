# _*_ coding: utf-8 _*_

'''
Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def diameterOfBinaryTree(self, root):
        """
        :type root: TreeNode
        :rtype: int
        需要找到两个深度最大的点且处于同一个root下
        """
        self.length = 0

        def deep(root):
            if not root:
                return 0
            left_D, right_D = deep(root.left), deep(root.right)
            self.length = max(self.length, left_D + right_D)
            return 1 + max(left_D, right_D)
        deep(root)
        return self.length
