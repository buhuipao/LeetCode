# _*_ coding: utf-8 _*_

'''
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. 
When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
'''


from collections import OrderedDict


class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self._len = capacity
        self._dict = OrderedDict([])

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if not self._dict.has_key(key):
            return -1
        self._dict[key] = self._dict.pop(key)
        return self._dict[key]

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        if self._dict.has_key(key):
            self._dict.pop(key)
        if len(self._dict) < self._len:
            self._dict.popitem(last=False)
        self._dict[key] = value
