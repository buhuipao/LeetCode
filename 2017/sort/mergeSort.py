# _*_ coding: utf-8 _*_


def mergeSort(_list):
    if len(_list) <= 1:
        return _list

    # 逐位对比，进行合并
    def merge(left, right):
        merged = []

        while left and right:
            merged.append(left.pop(0) if left[0] < right[0] else right.pop(0))

        while left:
            merged.append(left.pop(0))

        while right:
            merged.append(right.pop(0))

        return merged

    middle = int(len(_list) / 2)
    left = mergeSort(_list[:middle])
    right = mergeSort(_list[middle:])
    return merge(left, right)
