class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        # 统计频率
        freq = collections.Counter(tasks)

        # 找出最多执行次数
        maxExec = max(freq.values())

        # 统计同时具有最多运行次数的任务数
        maxCount = sum(1 for v in freq.values() if v == maxExec)
        # 此处+maxCount的理解：
        # 举例：n为2
        # A, A, A, B, B, B, C => A, (B), (C), A, (B), (), A, B
        # 如果出现了与最大频率相同的任务（例如 B），则会被追加到最后末尾（包括A自己）
        return max((maxExec-1)*(n+1)+maxCount, len(tasks))

