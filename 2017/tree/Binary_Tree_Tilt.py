# _*_ coding: utf-8 _*_

'''
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between \
the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
Note:

The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def findTilt(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root:
            return 0
        self.sum = 0
        self.helper(root)
        return self.sum

    def helper(self, root):
        if not root.left and not root.right:
            return root.val
        if not root.left:
            val = self.helper(root.right)
            self.sum += abs(val)
            return val + root.val
        if not root.right:
            val = self.helper(root.left)
            self.sum += abs(val)
            return val + root.val
        left_val = self.helper(root.left)
        right_val = self.helper(root.right)
        self.sum += abs(left_val - right_val)
        return left_val + right_val + root.val
