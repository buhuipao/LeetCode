# 最小栈左右遍历
class Solution:
    def longestMountain(self, A: List[int]) -> int:
        n = len(A)
        stack = []
        left, right, res = [0] * n, [0] * n, [0] * n

        for i in range(n):
            if stack and stack[-1] >= A[i]:
                stack = []
            if not stack or stack[-1] < A[i]:
                stack.append(A[i])
            left[i] = len(stack) - 1

        stack = []
        for i in range(n-1, -1, -1):
            if stack and stack[-1] >= A[i]:
                stack = []
            if not stack or stack[-1] < A[i]:
                stack.append(A[i])
            right[i] = len(stack) - 1

        for i in range(n):
            res[i] = left[i]+right[i]  + 1 if left[i] and right[i] else 0
        return max(res) if res else 0
