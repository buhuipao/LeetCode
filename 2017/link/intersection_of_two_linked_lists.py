# _*_ coding: utf-8 _*_

'''
Write a program to find the node at which the intersection of two singly linked lists begins.


For example, the following two linked lists:

A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗            
B:     b1 → b2 → b3
begin to intersect at node c1.

Notes:
If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        整体思路就是先截取较长的一部分，然后对齐后一起比较相等就好
        """
        def length(head):
            n = 0
            while head:
                n += 1
                head = head.next
            return n
        if not (headA and headB):
            return None
            
        A_len = length(headA)
        B_len = length(headB)
        if B_len > A_len:
            count = B_len - A_len
            while count > 0:
                headB = headB.next
                count -= 1
        else:
            count = A_len - B_len
            while count > 0:
                headA = headA.next
                count -= 1
        while headA and headB:
            if headA == headB:
                return headA
            headA, headB = headA.next, headB.next
        return None
