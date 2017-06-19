# _*_ coding: utf-8 _*_

'''
Given an integer, write a function to determine if it is a power of three.

Follow up:
Could you do it without using any loop / recursion?
'''


class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        import math
        if n < 1:
            return False
        # 四舍五入, 因为最后的取对数会有少许误差
        return 3 ** round(math.log(n, 3)) == n

    def isPowerOfThree1(self, n):
        """
        :type n: int
        :rtype: bool
        思路是在看了别人的解析后得出:
            直接在32位的有符号整型号里，找到最大的3的幂数，然后用n对它进行取余即可
            In [178]: math.log(1162261467, 2)
            Out[178]: 30.11428751370197

            In [180]: math.log(1162261467 * 3, 2)
            Out[180]: 31.699250014423125
            32位最大3的幂数是1162261467
        """
        return n > 0 and 1162261467 % n == 0
