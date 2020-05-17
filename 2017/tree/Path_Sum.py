# _*_ coding: utf-8 _*_

'''
Given a binary tree and a sum, determine if the tree has a root-to-leaf path \
such that adding up all the values along the path equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution1(object):
    def hasPathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: bool
        使用递归即可, 找到所有的正确路径即可
        """
        result = self.pathHelper(root, sum)
        return True if result else False

    def pathHelper(self, root, sum):
        if not root:
            return []
        val = root.val
        if not root.left and not root.right:
            if val == sum:
                return [[val]]
            else:
                return []
        l = self.pathHelper(root.left, sum-val)
        r = self.pathHelper(root.right, sum-val)
        result = []
        for i in l + r:
            result.append([val] + i)
        return result


class Solution(object):
    def hasPathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: bool
        使用DFS，每次迭代更深的层次，最后迭代到叶子节点验证和目标值是否相等
        """
        if not root:
            return False
        stack = []
        stack.append((root, root.val))
        while stack:
            curNode, curSum = stack.pop()
            if not curNode.left and not curNode.right and curSum == sum:
                return True
            if curNode.right:
                stack.append((curNode.right, curSum+curNode.right.val))
            if curNode.left:
                stack.append((curNode.left, curSum+curNode.left.val))
        return False

    def hasPathSum1(self, root, sum):
        """
        使用递归直接找出结果
        """
        if not root:
            return False
        if not root.left and not root.right and root.val == sum:
            return True
        return self.hasPathSum(root.left, sum-root.val) or \
            self.hasPathSum(root.right, sum-root.val)
