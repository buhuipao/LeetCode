# _*_ coding: utf-8 _*_

'''
Given a linked list and a value x, partition it such that all
nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
'''

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        if not head or not head.next:
            return head
        node = head
        while node and node.next:
            pre = node
            node = node.next.next
        # 确定老的尾巴
        if node:
            old_tail = node
        else:
            old_tail = pre.next
        # 确定新的头
        pre_head = ListNode(None)
        pre_head.next = head
        node = head
        while node:
            if node.val >= x:
                pre_head.next = node.next
                # 全部大于x
                if node == old_tail:
                    return head
                node = node.next
            else:
                break
        # 开始
        pre = ListNode(None)
        node = head
        tail = old_tail
        while node != old_tail:
            if node.val >= x:
                temp = node.next
                pre.next = temp
                tail.next = node
                node.next = None
                tail = tail.next
                node = temp
            else:
                pre = node
                node = node.next
        # 处理旧的尾巴
        if node.val >= x:
            tail.next = node
            temp = node.next
            pre.next = temp
            node.next = None
        return pre_head.next

    def partition1(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        h1 = l1 = ListNode(None)
        h2 = l2 = ListNode(None)
        while head:
            if head.val < x:
                l1.next = head
                l1 = l1.next
            else:
                l2.next = head
                l2 = l2.next
            head = head.next
        l1.next = h2.next
        l2.next = None
        return h1.next
