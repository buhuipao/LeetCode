# _*_ coding: utf-8 _*_

# 合并有序的A，B列表


class Merge:
    def mergeAB(self, A, B, n, m):
        for i in xrange(n):
            while B:
                if B[0] <= A[i]:
                    A.insert(i, B.pop(0))
                else:
                    break
        return A
