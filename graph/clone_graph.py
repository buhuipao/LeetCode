# _*_ coding: utf-8 _*_

'''
Clone an undirected graph. Each node in the graph contains a label and a list of its neighbors.


OJ's undirected graph serialization:
Nodes are labeled uniquely.

We use # as a separator for each node, and , as a separator for node label and each neighbor of the node.
As an example, consider the serialized graph {0,1,2#1,2#2,2}.

The graph has a total of three nodes, and therefore contains three parts as separated by #.

First node is labeled as 0. Connect node 0 to both nodes 1 and 2.
Second node is labeled as 1. Connect node 1 to node 2.
Third node is labeled as 2. Connect node 2 to node 2 (itself), thus forming a self-cycle.
Visually, the graph looks like the following:

       1
      / \
     /   \
    0 --- 2
         / \
         \_/
'''
import collections


class UndirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []


class Solution:
    # @param node, a undirected graph node
    # @return a undirected graph node
    def cloneGraph1(self, node):
        # BFS
        if not node:
            return
        
        _queue = collections.deque([node])
        record = {node.label: UndirectedGraphNode(node.label)}
        
        while _queue:
            item = _queue.popleft()
            for nei in item.neighbors:
                if nei.label not in record:
                    record[nei.label] = UndirectedGraphNode(nei.label)
                    _queue.append(nei)
                record[item.label].neighbors.append(record[nei.label])
        return record[node.label]
                    
    def cloneGraph(self, node):
        # DFS
        def dfs(node):
            if node.label in self.record:
                return self.record[node.label]
            
            self.record[node.label] = UndirectedGraphNode(node.label)
            for nei in node.neighbors:
                self.record[node.label].neighbors.append(dfs(nei))
            return self.record[node.label]
            
        if not node:
            return
        
        self.record = {}
        return dfs(node)
