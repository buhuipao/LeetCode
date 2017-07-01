# _*_ coding: utf-8 _*_

'''
Median is the middle value in an ordered integer list.
If the size of the list is even, there is no middle value.
So the median is the mean of the two middle value.
Examples:
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:

void addNum(int num) - Add a integer number from the data stream to the data structure.
double findMedian() - Return the median of all elements so far.
For example:

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
'''


class MedianFinder(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.list = []
        self.count = 0

    def addNum(self, num):
        """
        :type num: int
        :rtype: void
        二分查找
        """
        l, h = 0, self.count - 1
        mid = 0
        while l <= h:
            mid = (l + h) / 2
            if num <= self.list[mid]:
                h = mid - 1
            else:
                l = mid + 1
        self.list.insert(l, num)
        self.count += 1

    def findMedian(self):
        """
        :rtype: float
        """
        mid = self.count / 2
        if self.count % 2 == 1:
            return float(self.list[mid])
        else:
            return float(self.list[mid] + self.list[mid-1]) / 2

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
