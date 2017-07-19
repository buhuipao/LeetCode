# _*_ coding: utf-8 _*_

'''
You need to find the largest value in each row of a binary tree.

Example:
Input:

          1
         / \
        3   2
       / \   \
      5   3   9

Output: [1, 3, 9]
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def largestValues(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        按层遍历即可
        """
        if not root:
            return []
        queue, result = [root], []
        nlast = last = root
        temp = root.val
        while queue:
            node = queue.pop(0)
            if node.left:
                nlast = node.left
                queue.append(node.left)
            if node.right:
                nlast = node.right
                queue.append(node.right)
            temp = max(temp, node.val)
            if node == last:
                last = nlast
                result.append(temp)
                temp = float('-inf')
        return result
