# _*_ coding: utf-8 _*_

'''
A linked list is given such that each node contains an additional
random pointer which could point to any node in the list or null.

Return a deep copy of the list.
'''

# Definition for singly-linked list with a random pointer.
# class RandomListNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.next = None
#         self.random = None


class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        if not head:
            return None
        node = head
        while node:
            temp = node.next
            node.next = RandomListNode(node.label)
            node.next.next = temp
            node = temp
        # 新的头节点
        n_head = head.next
        node = head
        while node:
            if node.random:
                node.next.random = node.random.next
            else:
                node.next.random = None
            node = node.next.next
        node = head
        while node:
            temp = node.next
            node.next = temp.next
            # 判定是否为最后的节点
            if node.next:
                temp.next = node.next.next
            else:
                temp.next = None
            node = node.next
        return n_head
