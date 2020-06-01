from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        record = defaultdict(list)
        for s in strs:
            key = ''.join(sorted(s))
            record[key].append(s)

        return list(record.values())
