"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""
class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        m = {}
        node = head
        while node is not None:
            if node not in m:
                m[node] = Node(node.val)
            node = node.next

        for k in m:
            if k.next is not None:
                m[k].next = m[k.next]
            if k.random is not None:
                m[k].random = m[k.random]
        
        return m[head] if m else None
            
