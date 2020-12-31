# 尽可能的保留区间
class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        if not intervals:
            return 0
        intervals = sorted(intervals, key=lambda i: i[1])

        n = len(intervals)
        R = intervals[0][1] # 第一个右侧
        v = 1

        for i in range(1, n):
            if intervals[i][0] >= R: # 更新右侧边界
                v += 1
                R = intervals[i][1]

        return n - v
