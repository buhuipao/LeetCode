# _*_ coding: utf-8 _*_

'''
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes v and w as \
the lowest node in T that has both v and w as descendants \
(where we allow a node to be a descendant of itself).”

        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2       0       8
         /  \
         7   4
For example, the lowest common ancestor (LCA) of nodes 5 and 1 is 3.
Another example is LCA of nodes 5 and 4 is 5, \
since a node can be a descendant of itself according to the LCA definition.
'''

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        if not root:
            return None
        P = self.ancestor(root, p)
        Q = self.ancestor(root, q)
        h = min(len(P), len(Q))
        for i in xrange(h):
            if P[i] is Q[i]:
                result = P[i]
            else:
                break
        return result

    def ancestor(self, root, dis):
        """
        根据祖先找到目标节点
        """
        if not root:
            return []
        if root is dis:
            return [dis]
        l = self.ancestor(root.left, dis)
        r = self.ancestor(root.right, dis)
        if not l + r:
            return []
        return [root] + l + r
