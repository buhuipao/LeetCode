# _*_ coding: utf-8 _*_

'''
There is a brick wall in front of you. The wall is rectangular and has several rows of bricks.
The bricks have the same height but different width.
You want to draw a vertical line from the top to the bottom and cross the least bricks.

The brick wall is represented by a list of rows.
Each row is a list of integers representing the width of each brick in this row from left to right.

If your line go through the edge of a brick, then the brick is not considered as crossed.
You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.

You cannot draw a line just along one of the two vertical edges of the wall,
in which case the line will obviously cross no bricks.
'''


class Solution(object):
    def leastBricks(self, wall):
        """
        :type wall: List[List[int]]
        :rtype: int
        """
        width_dict = {}
        for row in wall:
            width = 0
            for i in xrange(len(row)-1):
                # 增加宽度
                width += row[i]
                if width not in width_dict:
                    width_dict[width] = 0
                width_dict[width] += 1
        if not width_dict:
            return len(wall)
        return len(wall) - max(width_dict.values())
