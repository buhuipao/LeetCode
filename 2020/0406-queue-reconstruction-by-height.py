class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        # 按照people[0]降序排列，同时people[1]升序排列
        # 降序排列可以使得每次的插入都忽略其身后的（已插入）的人
        people = sorted(people, key=lambda p: (-p[0], p[1]))
        ans = []
        for p in people:
            ans.insert(p[1], p)
        return ans
