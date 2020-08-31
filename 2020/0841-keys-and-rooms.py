class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        st = set()

        def dfs(n: int) -> None:
            if n in st:
                return 
            st.add(n)
            for i in rooms[n]:
                dfs(i)

        dfs(0)
        return len(rooms) == len(st)
