# 查找连通分量
class Solution:
    def numSimilarGroups(self, strs: List[str]) -> int:
        n = len(strs)
        uf = UF(n)
        def similar(A,B):
            n = len(A)
            cnt = 0
            for i in range(n):
                if A[i] != B[i]:
                    cnt += 1
            return cnt == 2
        for i in range(n):
            for j in range(i+1,n):
                # 两个字符串相同或者相识
                if strs[i] == strs[j] or similar(list(strs[i]),list(strs[j])):
                    uf.union(i,j)
        return uf.count

class UF:
    def __init__(self,n):
        self.parent=list(range(n))
        self.count=n
    def find(self,x):
        if x!=self.parent[x]:
            self.parent[x]=self.find(self.parent[x])
        return self.parent[x]
    def union(self,x,y):
        x,y=self.find(x),self.find(y)
        if x != y:
            self.parent[y] = x
            self.count -= 1
