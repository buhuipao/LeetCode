# _*_ coding: utf-8 _*_


'''
Given a binary tree, determine if it is height-balanced.
For this problem, a height-balanced binary tree is defined as a binary tree
in which the depth of the two subtrees of every node never differ by more than 1
'''
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        '''
        基本思路为：后序遍历，
        递归获取左右儿子的树的高度，如果左右儿子高度相差大于1返回Fasle, 否则返回树的高度为max(l_hight, r_hight)+1
        重复上步
        '''
        def dfs(node):
            # 空树也是平衡树, 返回高度0
            if not node:
                return 0
            l_hight = dfs(node.left)
            r_hight = dfs(node.right)
            if abs(l_hight - r_hight) > 1:
                self.is_balance = False
            # 返回树的高度
            return max(l_hight, r_hight) + 1

        self.is_balance = True
        dfs(root)
        return self.is_balance
