# _*_ coding: utf-8 _*_

'''
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.

'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return None
            
        pre_head = ListNode(None)
        pre_head.next = head
        node = ListNode(None)
        node.next = pre_head.next
        # 如：node->1->2->2.next, 每次交换1，2节点
        while node.next and node.next.next:
            temp = ListNode(None)
            # temp指向要交换的第二个点
            temp.next = node.next.next

            # 第一步：修改第一个交换点的next为交换的第二个点的next
            node.next.next = temp.next.next
            # 第二部：修改第二个交换点的next为第一个交换的点
            temp.next.next = node.next
            # 判断是否为第一次交换，必须保证头节点的正确
            if node.next == pre_head.next:
                pre_head.next = temp.next
            # 第三步：链接node到第二个交换点 
            node.next = temp.next
            
            # 修改新node节点为第一个交换点
            node = temp.next.next
            
        return pre_head.next
            
