class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        rows, cols = len(s) + 1, len(p) + 1
        # 初始化
        dp = [[False] * cols for _ in range(rows)]
        # init dp
        for j in range(cols):
            # 初始化dp[0][0]和p的任意个前缀*号匹配s为空字符串的情况
            if j == 0 or p[j - 1] == '*':
                dp[0][j] = True
            else:
                break
        # process dp
        for i in range(1, rows):
            for j in range(1, cols):
                # 如果p[j-1]为*，则消耗一个s的字符或者消耗*号自身
                if p[j - 1] == '*':
                    dp[i][j] = (dp[i - 1][j] | dp[i][j - 1])
                # 单个字符精确匹配或者替换匹配的情况
                elif p[j - 1] in {s[i - 1], '?'}:
                    dp[i][j] = dp[i - 1][j - 1]
        return dp[-1][-1]
