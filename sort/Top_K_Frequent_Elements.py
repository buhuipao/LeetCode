# _*_ coding: utf-8 _*_

'''
Given a non-empty array of integers, return the k most frequent elements.

For example,
Given [1,1,1,2,2,3] and k = 2, return [1,2].

Note: 
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
'''


class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """ 
        if len(nums) <= 1:
            return nums
            
        left = lambda i: 2 * i + 1
        right = lambda i: 2 * i + 2
        num_dict = {}
        for num in nums:
            num_dict[num] = num_dict.get(num, 0) + 1
        _nums = num_dict.items()
        def heap(_nums, i, k):
            while True:
                if left(i) < k and _nums[left(i)][1] < _nums[i][1]:
                    _min = left(i)
                else:
                    _min = i
                if right(i) < k and _nums[right(i)][1] < _nums[_min][1]:
                    _min = right(i)
                if _min == i:
                    break
                _nums[i], _nums[_min] = _nums[_min], _nums[i]
                i = _min
        # 建立调整小根堆
        for i in xrange(k-1, -1, -1):
            heap(_nums, i, k)
        
        for num in _nums[k:]:
            if num[1] > _nums[0][1]:
                _nums[0] = num
                heap(_nums, 0, k)
        return [i for [i, j] in _nums[:k]]
