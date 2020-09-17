# coding:utf-8

class UnionFind(object):
    def __init__(self, n):
        self.pre = list(range(n))

    def find(self, x):
        if x != self.pre[x]:
            self.pre[x] = self.find(self.pre[x])
        return self.pre[x]

    def union(self, x1, x2):
        root1 = self.find(x1)
        root2 = self.find(x2)
        if root1 != root2:
            self.pre[root2] = root1
            return False

        return True

class Solution(object):
    def findRedundantDirectedConnection(self, edges):
        """
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        n = len(edges)
        uf = UnionFind(n + 1)

        last = []
        parent = {}
        candidates = []
        for st, ed in edges:
            # 如果ed的入度已存在，那么证明ed入度为2，
            # 则在candidates同时有序添加这两个入度
            '''
              1
             / \
            v   v
            2-->3
            '''
            if ed in parent:
                candidates.append([parent[ed], ed])
                candidates.append([st, ed])
            else:
                # 记录入度：st -> ed
                parent[ed] = st
                ''' 如果st和ed已经联通并且还存在直接的连接关系，则替换最后一组环节点
                5 <- 1 -> 2
                     ^    |
                     |    v
                     4 <- 3
                '''
                if uf.union(st, ed):
                    last = [st, ed]

        '''
        分为两种情况：
        1) 都是入度为1， 则找出构成环的最后一条边；
        2)有入度为2的两条边(A->B, C->B)，则删除的边一定是在其中；
        先不将C->B加入并查集中，若不能构成环，则C->B是需要删除的点边，
        反之，则A->B是删除的边(去掉C->B还能构成环，则C->B一定不是要删除的边)
        '''
        if not candidates:
            return last
            
        return candidates[0] if last else candidates[1]
