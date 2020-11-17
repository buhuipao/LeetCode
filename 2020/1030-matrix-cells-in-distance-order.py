class Solution:
    def allCellsDistOrder(self, R: int, C: int, r0: int, c0: int) -> List[List[int]]:
        ans = []
        for i in range(R):
            for j in range(C):
                ans.append([i,j])  
        return sorted(ans, key=lambda i: abs(r0-i[0]) + abs(c0-i[1]))
