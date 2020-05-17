# _*_ coding: utf-8

'''
Follow up for problem "Populating Next Right Pointers in Each Node".

What if the given tree could be any binary tree? Would your previous solution still work?

Note:

You may only use constant extra space.
For example,
Given the following binary tree,
         1
       /  \
      2    3
     / \    \
    4   5    7
After calling your function, the tree should look like:
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \    \
    4-> 5 -> 7 -> NULL
添加一个右向指针
'''


class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None


class Solution:
    # @param root, a tree link node
    # @return nothing
    # O(logN)
    def connect(self, root):
        if not root:
            return root
        _queue = [root]
        n_last = last = root
        while _queue:
            node = _queue.pop(0)
            if node.left:
                _queue.append(node.left)
                n_last = node.left
            if node.right:
                _queue.append(node.right)
                n_last = node.right
            if node == last:
                node.next = None
                last = n_last
            else:
                node.next = _queue[0] if _queue else None
