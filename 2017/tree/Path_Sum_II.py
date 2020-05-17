# _*_ coding: utf-8 _*_

'''
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: List[List[int]]
        需要用递归找到叶子节点，然后就使处理如何拼接列表
        """
        if not root:
            return []
        val = root.val
        # 找到叶子节点就返回二维数组
        if not root.left and not root.right:
            if val == sum:
                return [[val]]
            else:
                return []
        l = self.pathSum(root.left, sum-root.val)
        r = self.pathSum(root.right, sum-root.val)
        result = []
        for i in l + r:
            result.append([val] + i)
        return result
