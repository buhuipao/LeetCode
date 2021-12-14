# 贪心：总是用当前课程去替换优先队列中已加入的最耗时课程，使得整体耗时降低，以便容纳下一个课程
class Solution:
    def scheduleCourse(self, courses: List[List[int]]) -> int:
        courses.sort(key=lambda c: c[1])

        q = list()
        # 优先队列中所有课程的总时间
        total = 0

        for ti, di in courses:
            if total + ti <= di:
                total += ti
                # Python 默认是小根堆
                heapq.heappush(q, -ti)
            elif q and -q[0] > ti:
                total -= -q[0] - ti
                heapq.heappop(q)
                heapq.heappush(q, -ti)

        return len(q)
