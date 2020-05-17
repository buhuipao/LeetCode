# _*_ coding: utf-8 _*_

'''
Given a singly linked list where elements are sorted in ascending order,
convert it to a height balanced BST.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def sortedListToBST(self, head):
        """
        :type head: ListNode
        :rtype: TreeNode
        每次找到中位数，然后递归
        """
        if not head or not head.next:
            return TreeNode(head.val) if head else head
        
        slow, fast, pre = head, head, None
        while fast and fast.next:
            pre, slow, fast = slow, slow.next, fast.next.next
        # 截断链表
        pre.next = None
        root = TreeNode(slow.val)
        # 递归查找中间位置的节点
        root.left, root.right = map(self.sortedListToBST, [head, slow.next])
        
        return root
