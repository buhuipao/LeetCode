# _*_ coding: utf-8 _*_

'''
Given a Binary Search Tree (BST), convert it to a Greater Tree such that \
every key of the original BST is changed to the original key plus sum of \
all keys greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def convertBST(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        把原BST节点的值变成原BST中比它大的值的与自己的和
        反中序遍历即可
        """
        self.sum = 0

        def travelTree(root):
            if not root:
                return
            _stack = [root]
            node = root.right
            while node or _stack:
                while node:
                    _stack.append(node)
                    node = node.right
                node = _stack.pop()
                self.sum += node.val
                node.val = self.sum
                node = node.left
        travelTree(root)
        return root
