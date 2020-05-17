# _*_ coding: utf-8 _*_

'''
Given preorder and inorder traversal of a tree, construct the binary tree.
Note:
You may assume that duplicates do not exist in the tree.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """
        if not preorder:
            return None
        root = TreeNode(preorder[0])
        mid = inorder.index(preorder[0])
        # 左子树位于preorder以及inorder的区间
        l = self.buildTree(preorder[1:mid+1], inorder[:mid])
        # 右子树位于的区间
        r = self.buildTree(preorder[mid+1:], inorder[mid+1:])
        root.left, root.right = l, r
        return root


class Solution1(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """
        if not preorder:
            return None
        root = TreeNode(preorder[0])
        stack = [root]
        i, j = 1, 0
        while i < len(preorder):
            temp = None
            cur = TreeNode(preorder[i])
            while stack and stack[-1].val == inorder[j]:
                temp = stack.pop()
                j += 1
            if temp:
                temp.right = cur
            else:
                stack[-1].left = cur
            stack.append(cur)
            i += 1
        return root
