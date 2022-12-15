class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        # 按照p[0]降序，p[1]升序排列
        people.sort(key=lambda p: (-p[0], p[1]))
        ans = []
        # 先把高个子排好，这样再往后排矮个子时就可以保证一件事：后续插入矮个子不会对队伍有影响；
        # 这个时候再去排列后续矮个子时，只需要将它插入到指定索引位置即可，因为之前插入的人都比他高
        for p in people:
            ans.insert(p[1], p)
        return ans
