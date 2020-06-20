class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        # 如果p为空，返回s是否为空即可
        if not p:
            return not s
        # 当前s、p的第一个值是否匹配
        cur_match = False
        if s:
            cur_match = p[0] in {'.', s[0]}
        # p中有＊号的情况
        if len(p) >= 2 and p[1] == '*':
            # 0个或者多个s[0]字符
            # 比如：(abc, b*abc) or (aabc, a*bc)
            return self.isMatch(s, p[2:]) or (cur_match and self.isMatch(s[1:], p))
        # 最普通的情况
        # 比如：（abc, abc) 或者 (abc, .bc)
        return cur_match and self.isMatch(s[1:], p[1:])
            
