class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if n == 1:
            return [0]
        
        g = [[] for _ in range(n)]
        for e in edges:
            src, des = e[0], e[1]
            g[src].append(des)
            g[des].append(src)
            
        parent = [0] * n
        maxL, node = 0, -1
        def dfs(x: int, p: int, l: int):
            nonlocal maxL, node
            if l > maxL:
                maxL, node = l, x
            parent[x] = p

            for y in g[x]:
                if y != p:
                    dfs(y, x, l+1)
        
        dfs(0, -1, 1)
        maxL = 0
        dfs(node, -1, 1)

        path = []
        while node != -1:
            path.append(node)
            node = parent[node]
        
        l = len(path)
        if l % 2 == 1:
            return path[l//2:l//2+1]

        return path[l//2-1:l//2+1]
