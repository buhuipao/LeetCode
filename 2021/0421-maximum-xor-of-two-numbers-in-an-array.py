class Trie:
    def __init__(self, val):
        self.val = val
        self.child = {}


class Solution:
    def findMaximumXOR(self, nums: List[int]) -> int:
        
        #取得最大长度
        L = len(format(max(nums), 'b'))-1

        # 构建前缀树
        root = Trie(-1)
       
        for n in nums:
            curr = root
            for i in range(L, -1, -1):

                v = (n >> i) & 1
                if v not in curr.child:
                    curr.child[v] = Trie(v)

                curr = curr.child[v]

        res = 0

        #搜索
        for n in nums:
            curr = root
            total = 0
            for i in range(L, -1, -1):
                v = (n >> i) & 1
                if 1-v in curr.child:
                    total = total * 2 + 1
                    curr = curr.child[1-v]
                else:
                    total = total * 2
                    curr = curr.child[v]
            
            res = max(res, total)

        return res
