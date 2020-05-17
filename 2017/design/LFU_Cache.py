# _*_ coding: utf-8 _*_

'''
Design and implement a data structure for Least Frequently Used (LFU) cache.
It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present.
When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item.
For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency),
the least recently used key would be evicted.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LFUCache cache = new LFUCache( 2 /* capacity */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.get(3);       // returns 3.
cache.put(4, 4);    // evicts key 1.
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
'''


class LFUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.data = {}
        self.counter = {}

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.data:
            return -1
        count = self.data[key]['used']
        self.data[key]['used'] += 1
        if count+1 not in self.counter:
            self.counter[count+1] = []
        self.counter[count].remove(key)
        self.counter[count+1].append(key)
        return self.data[key]['value']

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        if not self.capacity:
            return
        if key in self.data:
            self.data[key]['value'] = value
            count = self.data[key]['used']
            self.counter[count].remove(key)
            self.counter[count].append(key)
            return
        if len(self.data) < self.capacity:
            self.data[key] = {'value': value, 'used': 0}
            if 0 not in self.counter:
                self.counter[0] = []
            self.counter[0].append(key)
            return
        i = 0
        while True:
            if self.counter[i]:
                break
            else:
                i += 1
        old_key = self.counter[i].pop(0)
        self.data.pop(old_key)
        self.data[key] = {'value': value, 'used': 0}
        self.counter[0].append(key)

# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
