# coding: utf-8

'''
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
'''

import sys

def max_sum_subseq(L):
    This_Sum = 0
    Max_Sum = 0
    result = {}
    result['max'] = 0
    count = 0
    for i in xrange(len(L)):
        This_Sum += L[i]
         #如果此段和小于0了, 再加下去也会导致之后加进去和变小, 所以干脆置0
        if This_Sum < 0:
            This_Sum = 0
            count = 0
         #如果新的和大于之前最大值，立即替换
        elif This_Sum > Max_Sum:
            Max_Sum = This_Sum
            count += 1
             #跟随最大值一起变化，i存储此时的位置, count存储当前的累加项数,最后求队列是只需往前推count项即可
             #所以后面的else也是必须的, 没比Max_sum大时也要自增
            result['max'] = (i, count)
        else:
            count += 1
    #返回三个值, 分别是最大值, 总计前推项, 最大列表的结尾
    return (Max_Sum, result['max'][1], result['max'][0])

L = []
while True:
    line = sys.stdin.readline()
    if not line:
        break
    L.extend(line.strip('[\n ]').split(','))
L = map(int, L)
result, k, end = max_sum_subseq(L)
print result, k, L[end+1-k:end+1]
