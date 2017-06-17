# _*_ coding: utf-8 _*_

'''
Given a list, rotate the list to the right by k places, where k is non-negative.

For example:
Given 1->2->3->4->5->NULL and k = 2,
return 4->5->1->2->3->NULL.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not (head and head.next):
            return head
        
        length = 1
        node = head
        while node.next:
            node = node.next
            length += 1
        node.next = head
        
        tail, cur = node, head
        
        # 右旋k, 其实就是左旋(length - k % length)
        n = length - k % length
        while n > 0:
            cur = cur.next
            tail = tail.next
            n -= 1
        tail.next = None
        
        return cur
