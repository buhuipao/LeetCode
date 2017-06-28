# _*_ coding: utf-8 _*_

'''
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
Given n will always be valid.
Try to do this in one pass.
'''


# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        i = 0
        node = head
        while node:
            node = node.next
            i += 1
        count = i - n
        pre = ListNode(None)
        pre.next = head
        node = head
        while count:
            pre = node
            node = node.next
            count -= 1
        pre.next = node.next
        # 检查剔除的节点是否为头节点
        if node == head:
            return pre.next
        return head
