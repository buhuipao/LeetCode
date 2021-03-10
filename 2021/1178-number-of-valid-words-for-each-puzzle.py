class Solution:
    def findNumOfValidWords(self, words: List[str], puzzles: List[str]) -> List[int]:
        freq = collections.Counter()
        for word in words:
            mask = 0
            for c in word:
                mask |= 1 << (ord(c) - ord('a'))
            freq[mask] += 1
        res = []
        for puzzle in puzzles:
            total = 0
            for perm in self.subsets(puzzle[1:]):
                mask = 1 << (ord(puzzle[0]) - ord('a'))
                for c in perm:
                    mask |= 1 << (ord(c) - ord('a'))
                total += freq[mask]
            res.append(total)
        return res
    
    def subsets(self, words: List[int]) -> List[List[int]]:
        res = [""]
        for i in words:
            res = res + [i + word for word in res]
        return res
