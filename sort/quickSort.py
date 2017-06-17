import random


def qSort(a):
    if len(a) <= 1:
        return a
    else:
        # 基准的选择不同于前，是从数组中任意选择一个值做为基准
        q = random.choice(a)
        return qSort([elem for elem in a if elem < q]) +\
                [q] * a.count(q) +\
                qSort([elem for elem in a if elem > q])


class sort(object):
    def quickSort(self, A, n):
        less, pivot_list, more = [], [], []
        # write code here
        if n <= 1:
            return A
        else:
            pivot = A[0]
            for i in A:
                if i < pivot:
                    less.append(i)
                elif i > pivot:
                    more.append(i)
                else:
                    pivot_list.append(i)
        less = self.quickSort(self, less, len(less))
        more = self.quickSort(self, more, len(more))

        return less + pivot_list + more


# 方法5
# 这个最狠了，一句话搞定快速排序，瞠目结舌吧

qs = lambda xs: ((len(xs) <= 1 and [xs]) or [qs([x for x in xs[1:] if x < xs[0]]) + [xs[0]] + qs([x for x in xs[1:] if x >= xs[0]]) ] )[0]
