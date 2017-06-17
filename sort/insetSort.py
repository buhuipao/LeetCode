# _*_ coding: utf-8 _*_


# 时间复杂度，n^2
def insertionSort(A):
    n = len(A)
    for j in xrange(1, n):
        for i in xrange(j, 0, -1):
            if A[i] < A[i-1]:
                A[i], A[i-1] = A[i-1], A[i]
            else:
                break
    return A
