from sortedcontainers import SortedList

class Solution:
    def containsNearbyAlmostDuplicate(self, nums: List[int], k: int, t: int) -> bool:
        # O(N logk)
        window = SortedList()
        for i in range(len(nums)):
            # len(window) == k
            if i > k:
                window.remove(nums[i - 1 - k])
            window.add(nums[i])
            idx = bisect.bisect_left(window, nums[i])   # 找到当前值在有序列表中的index
            if idx > 0 and abs(window[idx] - window[idx-1]) <= t:   # 检查有序列表中index左边的一个位置是否符合条件 abs(nums[i] - nums[j]) <= t
                return True
            if idx < len(window) - 1 and abs(window[idx+1] - window[idx]) <= t: # 检查有序列表中index右边的一个位置是否符合条件 abs(nums[i] - nums[j]) <= t
                return True
        return False
