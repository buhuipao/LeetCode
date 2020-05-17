# _*_ coding: utf-8 _*_

'''
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
Note: You may assume k is always valid, 1 ? k ? BST's total elements.
Follow up: What if the BST is modified (insert/delete operations) often and
you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
'''


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def kthSmallest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        _stack = [root]
        node = root.left
        while _stack or node:
            while node:
                _stack.append(node)
                node = node.left
            node = _stack.pop(-1)
            k -= 1
            if k < 1:
                return node.val
            node = node.right
