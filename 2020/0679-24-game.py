'''
递归
[a, b, c, d]可以转化为:
    1) [(a+b), c, d];
    2) [(a-b), c, d];
    3) [(a*b), c, d];
    4) [(a/b), c, d];
'''
class Solution:
    def judgePoint24(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return math.isclose(nums[0], 24)
        # 1)使用全排列工具
        # 2)为了防止b为0，导致a/b异常，使用 b and a/b 代替
        return any(self.judgePoint24([x] + rest)
            for a, b, *rest in itertools.permutations(nums) for x in {a+b, a-b, a*b, b and a/b})
