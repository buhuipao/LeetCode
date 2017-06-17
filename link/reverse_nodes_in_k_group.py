# _*_ coding: utf-8 _*_

'''
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list.
If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
'''

class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        # 测试余下的链表是否长度大于k, 并且返回end节点
        def test_len(start, k):
            n = 0
            while start:
                if start:
                    n += 1
                    if n >= k:
                        return True, start
                    start = start.next
            return False, None
            
        if not head:
            return None
        pre_head = ListNode(None)
        pre_head.next = head
        start = head
        last_end = ListNode(None)
        # 返回是否可以继续反转的状态和end节点
        stat, end = test_len(start, k)
        while stat:
            '''
            连接上一个end到下一个新的start(即:现在的end)
            '''
            last_end.next = end
            temp = ListNode(None)
            temp.next = end.next
            self.reverse(start, end)
            '''
            最后一个end就是经过反转前的start
            并判断是否为第一个反转，记得更新头节点
            '''
            last_end = start
            if start == head:
                pre_head.next = end
            '''
            连接已反转和剩余部分
            '''
            start.next = temp.next
            start = start.next
            '''
            测试并找出下一个end
            '''
            stat, end = test_len(start, k)
        
        return pre_head.next
    
    def reverse(self, start, end):
        '''
        反转一段链表比且返回新的头和尾节点
        '''
        pre = ListNode(None)
        _cur = start
        while pre != end:
            temp = ListNode(None)
            temp.next = _cur.next
            _cur.next = pre
            pre = _cur
            _cur = temp.next
        return end, start
