class Solution:
    def numJewelsInStones(self, J: str, S: str) -> int:
        m = set(J)
        return len([True for s in S if s in m])
