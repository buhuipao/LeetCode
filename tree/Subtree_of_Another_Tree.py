# _*_ coding: utf-8 _*_

'''
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s.
A subtree of s is a tree consists of a node in s and all of this node's descendants.
The tree s could also be considered as a subtree of itself. Example 1:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
Given tree t:
   4
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s. Example 2:
Given tree s:

     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.

'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def isSubtree(self, s, t):
        """
        :type s: TreeNode
        :type t: TreeNode
        :rtype: bool
        先序遍历，然后递归匹配
        或者可以先序列化为字符串，然后kmp
        """
        if not s or not t:
            return False
        stack = [s]
        while stack:
            node = stack.pop()
            if self.helper(node, t):
                return True
            if node.right:
                stack.append(node.right)
            if node.left:
                stack.append(node.left)
        return False

    def helper(self, t1, t2):
        if not t2 and not t1:
            return True
        if not t1 or not t2 or t2.val != t1.val:
            return False
        return self.helper(t1.left, t2.left) and \
            self.helper(t1.right, t2.right)
