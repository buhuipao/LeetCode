# _*_ coding: utf-8 _*_

'''
Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
'''


class Interval(object):
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e


class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        if len(intervals) <= 1:
            return intervals
        # 根据区间的左边界排序
        _inters = sorted(intervals, key=lambda inter: inter.start)

        result = [_inters[0]]
        start = _inters[0].start
        end = _inters[0].end
        for i in xrange(1, len(_inters)):
            if end >= _inters[i].start:
                if end < _inters[i].end:
                    end = _inters[i].end
                    result[-1] = Interval(start, end)
                else:
                    result[-1] = Interval(start, end)
            else:
                start, end = _inters[i].start, _inters[i].end
                result.append(_inters[i])
            # start, end = result[-1].start, result[-1].end

        return result

