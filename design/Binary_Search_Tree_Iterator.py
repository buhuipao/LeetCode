# _*_ coding: utf-8 _*_

'''
Implement an iterator over a binary search tree (BST).
Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

Note: next() and hasNext() should run in average O(1) time and uses O(h) memory,
where h is the height of the tree.
'''

# Definition for a  binary tree node
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class BSTIterator1(object):
    def __init__(self, root):
        """
        :type root: TreeNode
        BST的中序遍历
        """
        self.cur_node = root.left if root else None
        self.stack = [root] if root else []

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.stack or self.cur_node

    def next(self):
        """
        :rtype: int
        """
        while self.cur_node:
            self.stack.append(self.cur_node)
            self.cur_node = self.cur_node.left
        self.cur_node = self.stack.pop()
        val = self.cur_node.val
        self.cur_node = self.cur_node.right
        return val


class BSTIterator(object):
    def __init__(self, root):
        """
        :type root: TreeNode
        BST的中序遍历, 建立一个hash表
        """
        self.dict = {}
        cur_node = root.left if root else None
        stack = [root] if root else []
        _min = float('-inf')
        while stack or cur_node:
            while cur_node:
                stack.append(cur_node)
                cur_node = cur_node.left
            cur_node = stack.pop()
            val = cur_node.val
            self.dict[_min] = val
            _min = val
            cur_node = cur_node.right
        if not self.dict:
            self.max = float('-inf')
        else:
            self.max = val
        self.val = float('-inf')

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.val < self.max

    def next(self):
        """
        :rtype: int
        """
        result = self.dict[self.val]
        self.val = result
        return result

# Your BSTIterator will be called like this:
# i, v = BSTIterator(root), []
# while i.hasNext(): v.append(i.next())
