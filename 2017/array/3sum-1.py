class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums = self.fast_sort(nums)
        # 过滤全是正数全是负数以及len(nums)<3的情况
        if len(nums) < 3 or nums[0] > 0 or nums[-1] < 0:
            return []
        zero = filter(lambda x: x==0, nums)
        # 全零状态
        if zero == nums:
            return [[0, 0, 0]]
        over = filter(lambda x: x>=0, nums)
        less = filter(lambda x: x<0 , nums)
        result = []
        if len(less) <= len(over):
            for i in xrange(len(less)):
                if i != 0 and less[i-1] == less[i]:
                    continue
                j, k = i+1, len(nums)-1
                while j < k:
                    if nums[j] == nums[j+1]:
                        j += 1
                    if nums[k] == nums[k-1]:
                        k -= 1
                    if less[i] + nums[j] + nums[k] < 0:
                        j += 1
                    elif less[i] + nums[j] + nums[k] >0:
                        k -= 1
                    else:
                        result.append([less[i], nums[j], nums[k]])
                        j, k = j+1, k-1
                i += 1

        else:
            for i in xrange(len(over)-1, -1, -1):
                if i != len(over)-1 and over[i] == over[i+1]:
                    continue
                j, k = i-1+len(less), 0
                while j > k:
                    if nums[j] == nums[j-1]:
                        j -= 1
                    if nums[k] == nums[k+1]:
                        k += 1
                    if over[i] + nums[j] + nums[k] < 0:
                        k += 1
                    elif over[i] + nums[j] + nums[k] >0:
                        j -= 1
                    else:
                        result.append([over[i], nums[j], nums[k]])
                        j, k = j-1, k+1
                i += 1
                
        return result
    # 快排    
    def fast_sort(self, nums):
        if len(nums) < 2:
            return nums
        _num = nums[0]
        _gt, _le = [], []
        for num in nums[1:]:
            if num > _num:
                _gt.append(num)
            else:
                _le.append(num)
        return self.fast_sort(_le) + [_num] +  self.fast_sort(_gt)
