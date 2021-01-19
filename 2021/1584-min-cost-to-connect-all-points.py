class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        m = []
        n = len(points)
        # 第一步：计算所有点之间的距离
        for i in range(n):
            for j in range(i+1, n):
                m.append([i, j, abs(points[i][0]-points[j][0])+abs(points[i][1]-points[j][1])])
        # 以点与点之间的曼哈顿距离进行排序
        m = sorted(m, key=lambda i: i[2])

        # 第二步：定义并查集，用于判断边是否完全加入了最小生成树中
        parent = list(range(n))    # 初始化所有点的root为自己
        def find(x) -> int:
            if x != parent[x]:
                parent[x] = find(parent[x])
            return parent[x]
        def union(x, y) -> bool:
            a, b = find(x), find(y)
            if a != b:
                parent[a] = b
                return True # 新联结的
            return False

        # 第三步：进行加边
        ans, edges = 0, 0
        for i, j, d in m:
            # 如果两个点构成的边还没加入最小生成树中
            if union(i, j):
                ans += d
                edges += 1
            # 证明n个点已经加入了
            if edges == n-1:
                break
        return ans
        
