class Solution:
    def scheduleCourse(self, courses: List[List[int]]) -> int:
        # 基于deadline排序
        courses.sort(key=lambda c: c[1])

        q = list()
        cur_time = 0

        for t, dd in courses:
            if cur_time + t <= dd:
                cur_time += t
                # python默认是小根堆，所以需要反转
                heapq.heappush(q, -t)
            # 贪心：总是用当前课程去替换优先队列中已加入的最耗时课程，使得整体耗时降低，以便容纳下一个课程
            elif q and -q[0] > t:
                cur_time -= (-q[0] - t)
                heapq.heappop(q)
                heapq.heappush(q, -t)

        return len(q)
