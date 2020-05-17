# _*_ coding: utf-8 _*_

'''
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        if not head:
            return None
                
        pre_head = ListNode(None)
        pre_head.next = head
        pre_node = pre_head

        while head:
            if head.val == val:
                pre_node.next = head.next
            else:
                pre_node = head
            head = head.next

        return pre_head.next
