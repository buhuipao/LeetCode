class Solution:
    def kClosest(self, points: List[List[int]], K: int) -> List[List[int]]:
        return sorted(points, key=lambda p: (p[0]**2 + p[1]**2))[:K]
