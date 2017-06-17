# _*_ coding: utf-8 _*_

'''
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one
sorted array.
Note:
You may assume that nums1 has enough space (size that is greater or equal to
m + n) to hold additional elements from nums2. The number of elements
initialized in nums1 and nums2 are m and n respectively.
'''


class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: void Do not return anything, modify nums1 in-place instead.
        """
        while m > 0 and n > 0:
            # 假如nums1的数大于nums2的数，就向后搬运num1的数，m减1
            if nums1[m - 1] > nums2[n - 1]:
                nums1[m + n - 1] = nums1[m - 1]
                m -= 1
            # 否则就搬运num2的值到m+n-1位置，并且n减一
            else:
                nums1[m + n - 1] = nums2[n - 1]
                n -= 1
        # 有两种情况:
        #   假如nums2先用完，以上循环由n=0退出, 那么num1[:n], nums[:n]都为[]
        #   假如num1先用完，m=0退出，那么下步就是直接把num2[:n]整段拌匀至num1最前面, 最后的列表还是有序
        #   同时用完，m==n==0同第一种
        nums1[:n] = nums2[:n]
