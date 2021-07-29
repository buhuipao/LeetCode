class Solution:
    def pathInZigZagTree(self, label: int) -> List[int]:
        row = int(log2(label))+1
        ans = [0]*row
        
        while row:
            ans[row-1] = label
            label = (2**row-1-label+2**(row-1))//2
            row -= 1

        return ans
