# 暴力解法+分类讨论，非KMP
class Solution:
    def repeatedStringMatch(self, a: str, b: str) -> int:
        aS, bS = set(a), set(b)
        if aS & bS != bS: return -1
        
        times = math.ceil(len(b) / len(a))
        if (b in a*times): return times
        elif (b in a*(times+1)): return times+1
        else: return -1
