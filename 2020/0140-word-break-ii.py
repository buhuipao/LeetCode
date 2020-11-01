class Solution:
    # 带备忘录的记忆化搜索
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        res = []
        memo = [1] * (len(s)+1)
        wordDict = set(wordDict)

        def dfs(wordDict,temp,pos):
            # 回溯前先记下答案中有多少个元素
            num = len(res)                  
            if pos == len(s):
                res.append(" ".join(temp))
                return
            for i in range(pos,len(s)+1):
                # 添加备忘录的判断条件，原理
                # 1）初始状态，假设s[pos:]都是可以分割的；
                # 2）如果被之前的dfs修改成了false（不能分割），也就无需继续往下了；
                if memo[i] and s[pos:i] in wordDict:
                    temp.append(s[pos:i])
                    dfs(wordDict,temp,i)
                    temp.pop()
            # 答案中的元素没有增加，说明s[pos:]不能分割，修改备忘录       
            memo[pos] = 1 if len(res) > num else 0 
            
        dfs(wordDict,[],0)
        return res
