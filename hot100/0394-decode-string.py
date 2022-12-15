class Solution:
    def decodeString(self, s: str) -> str:
        stack, res, repeat = [], "", 0
        for i in s:
            if i == '[':
                stack.append([repeat, res])     # 保存当前的repeat和上次的结余
                res, repeat = "", 0
            elif i == ']':
                cur_preat, last_res = stack.pop()   # 重新计算res
                res = last_res + cur_preat * res
            elif '0' <= i <= '9':
                repeat = repeat * 10 + int(i)   # repeat可能大于10
            else:
                res += i
        return res
