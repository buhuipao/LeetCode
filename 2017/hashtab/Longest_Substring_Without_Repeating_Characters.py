# _*_ coding: utf-8 _*_

'''
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring,
"pwke" is a subsequence and not a substring.
'''


class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        result = 1
        index = {}
        pw_index = -1
        temp = 0
        for i in xrange(len(s)):
            if s[i] not in index:
                index[s[i]] = i
                # 防止全是不重复的字符，必须用temp记录经过的次数
                temp += 0
            else:
                # 更新前一个不重复子字串的索引位置
                if index[s[i]] > pw_index:
                    pw_index = index[s[i]]
                # 重置temp
                temp = 0
                # 记录字符前一次出现位置
                index[s[i]] = i
            result = max(result, temp, i-pw_index)
        return result
