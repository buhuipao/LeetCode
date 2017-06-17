# _*_ coding: utf-8 _*_

'''
Sort a linked list using insertion sort.
'''


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def insertionSortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        # 插入排序链表
        if not (head and head.next):
            return head

        # 定义初始节点
        sorted_tail, cur = head, head.next
        pre_head = ListNode(None)
        pre_head.next = head
        while sorted_tail.next:
            # 定义一个next指向head的空节点
            pre_node = ListNode(None)
            pre_node.next = pre_head.next
            if sorted_tail.val <= cur.val:
                # 如果有序直接往后走
                sorted_tail, cur = cur, cur.next
            else:
                # 从头开始比较直到找到不大于的点就停下准备插入
                while pre_node.next.val < cur.val:
                    pre_node = pre_node.next
                '''
                接下来要干两件事，
                第一抽出当前节点，第二插入到pre_node和pre_node.next之间
                最后更新cur, sorted_tail是没变的无需修改
                '''
                # 先把排序好的尾巴的next指向下一个
                sorted_tail.next = cur.next
                # 然后把cur指向它不大于的节点(pre_node.next)
                cur.next = pre_node.next
                # 最后把pre_node的next指向cur就好
                pre_node.next = cur
                # 如果cur.next为原来的老head，那么证明是在头节点插入，必须更新头节点
                if cur.next == pre_head.next:
                    pre_head.next = pre_node.next
                # 重新赋值cur
                cur = sorted_tail.next
        return pre_head.next
