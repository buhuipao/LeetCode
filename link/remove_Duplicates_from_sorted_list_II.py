# _*_ coding: utf-8 _*_

'''
Given a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head or not head.next:
            return head
        
        dummy = ListNode(None)
        node, dummy.next, last_uniq = head, head, dummy
        last_val = None
        while node:
            # 先找出一串相等值
            while node and last_val == node.val:
                node = node.next
            # 假如是最后一个值，那么直接break
            if not node or not node.next:
                last_uniq.next = node
                break
            # 更新最后的值
            last_val = node.val
            # 重新判断更新last_val后当前节点和下一个节点
            if node.val != node.next.val:
                last_uniq.next = node
                last_uniq = last_uniq.next
                node = node.next
        return dummy.next


# 以下为他人的解法
class Solution1(object):
    '''算法思路：
    每走一步判断当前节点是否应该被添加到列表里边
    '''
    def deleteDuplicates(self, head):
        dummy = tail = ListNode(None)
        while head:
            node = head
            # 这里他多判断了一次, 但是很灵活地跳出了所有重复的点
            if head.next and head.next.val == node.val:
                while head and head.val == node.val:
                    head = head.next
            else:
                tail.next, tail, head = head, head, head.next

        tail.next = None
        return dummy.next
