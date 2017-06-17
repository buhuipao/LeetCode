# _*_ coding: utf-8 _*_


class BubbleSort:
    def bubbleSort(self, A, n):
        for i in xrange(1, n-1):
            for j in xrange(0, n-i):
                if A[j] > A[j+1]:
                    A[j], A[j+1] = A[j+1], A[j]
        return A
