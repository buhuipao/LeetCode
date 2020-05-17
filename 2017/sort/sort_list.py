# _*_ coding: utf-8 _*_

'''
Sort a linked list in O(n log n) time using constant space complexity.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    """
    :type head: ListNode
    :rtype: ListNode
    """
    def merge(self, h1, h2):
        '''
        dummy = tail = ListNode(None)
        while h1 and h2:
            if h1.val < h2.val:
                tail.next, tail, h1 = h1, h1, h1.next
            else:
                tail.next, tail, h2 = h2, h2, h2.next

        tail.next = h1 or h2
        return dummy.next
        '''
        # 先申请出头节点的前节点，以及尾节点
        pre_head = tail = ListNode(None)
        # 选举头节点, pre_head.next就是head的意思
        if h1 and h2:
            if h1.val < h2.val:
                pre_head.next = h1
            else:
                pre_head.next = h2
        # 假设一个节点其实为空，那么pre_head.next就是不为空的节点
        else:
            pre_head.next = h1 or h2

        while h1 and h2:
            # 首先把尾节点的next指到比较下更小的节点，
            # 然后往前移动尾节点更新tail，然后往前移动更新h1/h2
            if h1.val < h2.val:
                tail.next = h1
                tail = h1
                h1 = h1.next
            else:
                tail.next = h2
                tail = h2
                h2 = h2.next

        # 选择尚不为空的链作为尾巴的下一个指向, 其实就是串接
        tail.next = h1 or h2
        # 返回头节点
        return pre_head.next

    def sortList(self, head):
        # 空链表或者单节点链表
        if not head or not head.next:
            return head

        pre, slow, fast = None, head, head
        while fast and fast.next:
            # fast的步长为2，pre／slow的步长为1，pre紧挨着slow，
            # 最终状态为：fast为末尾节点或者为空, slow为中间节点
            # slow把链表拆为近似均等两半，并归排序
            pre, slow, fast = slow, slow.next, fast.next.next
        # pre为slow前一个节点，指向空节点就是拆分链表,
        # pre为第一段的尾节点，fast为后一段的尾节点(链表长度为2*n+1)或者为尾节点后一个(链表长度为2*n)
        pre.next = None

        # 合并以head和slow(中间节点)作为头节点的链表
        return self.merge(self.sortList(head), self.sortList(slow))
