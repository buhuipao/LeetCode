class Solution:
    def totalNQueens(self, n: int) -> int:
        col=set()
        pos=set()
        neg=set()
        res=0
        '''
        思路：
            正对角线（从左上到右下）的特点：横纵坐标之差相同
            副对角线（从右上到左下）的特点：横纵坐标之和相同
            每行、每列及正副对角线上只能出现一个皇后，由于从上到下放置皇后，所以只需考虑当前坐标对应的列和正副对角线是否已被访问即可
        作者：yim-6
        链接：https://leetcode-cn.com/problems/n-queens-ii/solution/python3-san-chong-fang-fa-shi-xian-nhuang-hou-ii-b/
        '''
        def backtrack(i):
            nonlocal res
            if i==n:
                res+=1
                return 
            for j in range(n):
                if j not in col and i-j not in pos and i+j not in neg:
                    #做选择
                    col.add(j)
                    pos.add(i-j)
                    neg.add(i+j)
                    #递归进入下一行
                    backtrack(i+1)
                    #撤销选择
                    col.remove(j)
                    pos.remove(i-j)
                    neg.remove(i+j)
        
        backtrack(0)
        return res
