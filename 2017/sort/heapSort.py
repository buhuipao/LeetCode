# _*_ coding: utf-8 _*_


def heapSort(_list):
    left = lambda i: i * 2 + 1
    right = lambda i: i * 2 + 2

    # 维护i位置
    def v_heap(_list, i):
        print(heap_size)
        while True:
            # 先找出i，left(i), right(i)中最大的，再判断i是否等于_max
            if left(i) < heap_size and _list[left(i)] > _list[i]:
                _max = left(i)
            else:
                _max = i
            if right(i) < heap_size and _list[right(i)] > _list[_max]:
                _max = right(i)
            # 成立则证明i的位置不需要调整, 否则就需要调整并沿着调整的路线往下调整
            if _max == i:
                break
            # 交换之后给i重新赋值
            else:
                _list[_max], _list[i] = _list[i], _list[_max]
                i = _max

    # 大根堆, 先从末尾往上层开始建立堆，每次都认为i之后的都是建立好的堆
    def build_heap(_list):
        for i in xrange(len(_list)-1, -1, -1):
            v_heap(_list, i)

    # 堆排序
    heap_size = len(_list)
    build_heap(_list)
    for i in xrange(len(_list)-1, 0, -1):
        _list[0], _list[i] = _list[i], _list[0]
        heap_size = i
        v_heap(_list, 0)
