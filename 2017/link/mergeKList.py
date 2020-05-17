# _*_ coding: utf-8 _*_

'''
Merge k sorted linked lists and return it as one sorted list.
Analyze and describe its complexity.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        采用并归的思想
        """
        # 过滤空链表
        _list = filter(lambda node: node, lists)
        return self.merge_help(_list)
        
    def merge_help(self, _list):
        n = len(_list)
        if n == 0:
            return None
        elif n == 1:
            return _list[0]
        elif n == 2:
            return self.merge(_list[0], _list[1])
        else:
            mid = n / 2
        
        return self.merge(self.merge_help(_list[:mid]), self.merge_help(_list[mid:]))

    def merge(self, l1, l2):
        if not (l1 and l2):
            return l1 or l2
            
        pre_head = ListNode(None)
        # 确定头节点
        pre_head.next = l1 if l1.val <= l2.val else l2
        node = ListNode(None)
        while l1 and l2:
            if l1.val <= l2.val:
                node.next = l1
                l1 = l1.next
            else:
                node.next = l2
                l2 = l2.next
            node = node.next
        node.next = l1 or l2
        return pre_head.next
