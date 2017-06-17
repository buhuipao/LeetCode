# _*_ coding: utf-8 _*_


# shellSort其实就是选取不同步长，然后进行多次插入排序
# 因为插入排序的时间复杂度和原始数据有序性有关，nlogn
def shellSort(A):
    n = len(A)
    step = 0
    # 选定最大步长，1, 4, 13, 40, 121...
    while step <= n/3:
        step = step * 3 + 1

    while step >= 1:
        for j in xrange(1, n):
            for i in xrange(j, 0, -step):
                if A[i] < A[i-step]:
                    A[i], A[i-step] = A[i-step], A[i]
                else:
                    break
        step /= 3
    return A
