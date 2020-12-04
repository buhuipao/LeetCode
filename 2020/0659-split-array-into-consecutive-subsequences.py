from collections import Counter, defaultdict
class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        dic = Counter(nums)
        tail = defaultdict(int)
        for n in nums:
            if not dic[n]:
                continue
            dic[n] -= 1
            if not tail[n-1]:
                if dic[n+1] and dic[n+2]:
                    dic[n+1] -= 1
                    dic[n+2] -= 1
                    tail[n+2] += 1
                else:
                    return False
            elif tail[n-1]:
                tail[n-1] -= 1
                tail[n] += 1
        return True 
