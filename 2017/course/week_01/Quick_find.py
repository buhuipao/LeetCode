# _*_ coding: utf-8 _*_

'''
联通集的快速实现
'''


class Quick_Find(object):
    def __init__(self, N):
        self._list = []
        for i in xrange(N):
            self._list[i] = i

    def connected(self, p, q):
        return self._list[p] == self._list[q]

    def union(self, x, y):
        X = self._list[x]
        Y = self._list[y]
        for i in xrange(len(self._list)):
            if self._list[i] == X:
                self._list[i] == Y


class Qick_Union(object):
    def __init__(self, N):
        self._list = []
        for i in xrange(N):
            self._list[i] = i

    def root(self, x):
        while self._list[x] != x:
            x = self._list[x]
        return x

    def union(self, x, y):
        # 这样的合并将会导致树的高度过高，最后find还是很缓慢
        # 可以考虑用多叉树来减缓树的高度
        self._list[self.root(x)] = self.root(y)

    def connected(self, q, p):
        return self.root(q) == self.root(p)


class Qick_Union_Wight(object):
    def __init__(self, N):
        self._list = []
        self._w = []
        for i in xrange(N):
            self._list[i] = i
            self._w[i] = 1

    def root(self, x):
        while self._list[x] != x:
            # 为了让树变得更平，所以我们可以在找到祖父后接它变为直系父节点
            self._list[x] = self._list[self._list[x]]
            x = self._list[x]
        return x

    def union(self, x, y):
        # 这里使用带有权重的比较方法, 总是让权重小的树成为大树的子树
        root_x = self.root(x)
        root_y = self.root(y)
        if root_x == root_y:
            return
        # 比较权重，并更新权重
        if self._w[root_x] > self._w[root_y]:
            self._list[root_y] = root_x
            self._w[root_x] += self._w[root_y]
        else:
            self._list[root_x] = root_y
            self._w[root_y] += self._w[root_x]

    def connected(self, q, p):
        return self.root(q) == self.root(p)
