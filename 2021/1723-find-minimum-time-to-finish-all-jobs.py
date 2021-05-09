class Solution:
    def minimumTimeRequired(self, jobs: List[int], k: int) -> int:

        def check(limit):
            # 剪枝：排序后，大的先拿出来试，如果方案不行，失败得更快
            arr = sorted(jobs)

            groups = [0] * k
            # 分成K 组，看看在这个limit 下 能不能安排完工作
            if backtrace(arr, groups, limit):
                return True
            else:
                return False


        def backtrace(arr, groups, limit):
            # 尝试每种可能性
            #print(arr, groups, limit)
            if not arr: return True #分完，则方案可行
            v = arr.pop()

            for i in range(len(groups)):
                if groups[i] + v <= limit:
                    groups[i] += v
                    if backtrace(arr, groups, limit):
                        return True
                    groups[i] -= v

                    # 剪枝，如果这个工人没分到活，那别人肯定得多干活了，那最后的结果必然不是最小的最大值，就不用继续试了。 
                    if groups[i] ==0:
                        break

            arr.append(v)

            return False
        
        #每个人承担的工作的上限，最小为，job 里面的最大值，最大为 jobs 之和
        l, r  = max(jobs), sum(jobs)

        while l < r:
            mid = (l + r)//2

            if check(mid):
                r = mid
            else:
                l = mid + 1

        return l
