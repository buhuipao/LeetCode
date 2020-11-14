class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        l = len(arr2)
        m = dict((v, i) for (i, v) in enumerate(arr2))
        # 关键在于如何计算没出现在arr1里的数字如何排序
        # 以len(arr2)+v作为依据即可
        return sorted(arr1, key=lambda v: m[v] if v in m else l + v)
