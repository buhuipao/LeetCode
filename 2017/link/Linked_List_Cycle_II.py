# _*_ coding: utf-8 _*_

'''
Given a linked list, return the node where the cycle begins.
If there is no cycle, return null.
Note: Do not modify the linked list.

Follow up:
Can you solve it without using extra space?
'''


# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        找到一个链表的循环开始节点
        """
        if not head:
            return
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if fast == slow:
                fast = head
                while fast != slow:
                    fast, slow = fast.next, slow.next
                return slow
        return
