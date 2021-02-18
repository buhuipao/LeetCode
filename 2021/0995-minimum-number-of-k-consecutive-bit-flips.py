class Solution:
    def minKBitFlips(self, A, K):
        windowFlip = res = 0
        for i in range(len(A)):
            if i >= K and A[i - K] == 2:
                windowFlip -= 1
            if (windowFlip % 2) == A[i]:
                if i + K > len(A):
                    return -1
                A[i] = 2
                windowFlip, res = windowFlip + 1, res + 1
        return res
