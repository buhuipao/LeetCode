class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        if len(points) == 0:
            return 0
        # 排序
        points = sorted(points, key=lambda i: i[0])
        # 合并
        ans = 1
        pre = points[0]
        n, i = len(points), 0
        while i < n:
            p = points[i]
            # 找出两点之间的公共区间
            x, y = max(p[0], pre[0]), min(p[1], pre[1])
            if x <= y:
                pre = [x, y]
                i += 1
            else:
                ans += 1
                pre = p
        return ans 
