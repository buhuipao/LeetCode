# _*_ coding: utf-8 _*_

'''
Time:  O(nlogn)
Space: O(n)

You are given an integer array nums and you have to
return a new counts array. The counts array has the
property where counts[i] is the number of smaller
elements to the right of nums[i].

Example:

Given nums = [5, 2, 6, 1]

To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
Return the array [2, 1, 1, 0].
查找后面有多少各点小于当前值
'''


class Solution(object):
    def countSmaller(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        res = [0] * len(nums)
        bst = self.BST()
        '''
        从后面录入节点是为了少一个循环，
        而且从后往前的纪录才是有效的，例如：[8,5,4,2,3,6]
        1) 从前往后录的将会最后两个点将会出错, 错误结果为：[5,3,2,0,1,4]
        2) 从后往前的录入结果为：[5,3,2,0,0,0]
        '''
        for i in xrange(len(nums)-1, -1, -1):
            bst.insertNode(nums[i])
            res[i] = bst.query(nums[i])
        return res

    class BST(object):
        class BSTreeNode(object):
            def __init__(self, val):
                self.val = val
                # self.count 用于统计在录入过程中，从自己身上踩过且比自己小的节点数量
                # 并不是什么统计自己处于BST里面的位置
                self.count = 0
                self.left = self.right = None

        def __init__(self):
            self.root = None

        def insertNode(self, val):
            '''
            基本思路是，每次插入节点都需要更新她所踩过的节点的self.count
            '''
            node = self.BSTreeNode(val)
            # 初始化头节点
            if not self.root:
                self.root = node
                return
            curr = self.root
            while curr:
                if node.val < curr.val:
                    # 假如一个节点的值比当前节点值小，那么就从左边下去，并更新此节点的self.count(踩过数量)
                    # 如果此节点存在左节点, 那么就递归替换当前节点, 不存在就作为当前节点左儿子
                    curr.count += 1
                    if curr.left:
                        curr = curr.left
                    else:
                        curr.left = node
                        break
                # 如果比当前节点值大，如果那么就递归替换为右儿子，且不更新踩过数量self.count
                else:
                    if curr.right:
                        curr = curr.right
                    else:
                        curr.right = node
                        break

        def query(self, val):
            count = 0
            curr = self.root
            while curr:
                # 如果val小于当前值，就直接递归查找左儿子, 不需要采集当前点上self.count(被踩过的次数)
                if val < curr.val:
                    curr = curr.left
                # 右转的时候表示val大于当前点，所以需要收集当前点上的self.count(被踩过的次数),
                # 并且加1，因为这也算在当前点上踩过
                elif val > curr.val:
                    count += 1 + curr.count
                    curr = curr.right
                else:
                    # 如果相等证明找到此点，也需要收集self.count(被踩过的次数)
                    return count + curr.count
            return 0
