class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        nums = {"1", "2", "3", "4", "5", "6", "7", "8", "9"}
        row = [set() for _ in range(9)]
        col = [set() for _ in range(9)]
        palace = [[set() for _ in range(3)] for _ in range(3)]  # 3*3
        blank = []

        # 初始化，按照行、列、宫 分别存入哈希表
        for i in range(9):
            for j in range(9):
                ch = board[i][j]
                if ch == ".":
                    blank.append((i, j))
                else:
                    row[i].add(ch)
                    col[j].add(ch)
                    palace[i//3][j//3].add(ch)

        def dfs(n):
            # 当所有空白被填满时，返回成功
            if n == len(blank):
                return True
            # 取出第n个空白点的横坐标和纵坐标
            i, j = blank[n]
            # 找出此点可选供选择的数字，即：9个数字减去纵&横&方块已有的数字
            rst = nums - row[i] - col[j] - palace[i//3][j//3]  # 剩余的数字
            # 没有可选则返回失败，此路不通
            if not rst:
                return False
            #  在可选的数字里进行回溯
            for num in rst:
                # 作出选择
                board[i][j] = num
                row[i].add(num)
                col[j].add(num)
                palace[i//3][j//3].add(num)
                # 下一层回溯
                if dfs(n+1):
                    return True
                # 撤销选择
                row[i].remove(num)
                col[j].remove(num)
                palace[i//3][j//3].remove(num)

        dfs(0)
