"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""

class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return None
        last, nlast = root, None
        pre_node = Node(next=root)
        queue = [root]
        while queue:
            node = queue.pop(0)
            if node.left:
                queue.append(node.left)
                nlast = node.left
            if node.right:
                queue.append(node.right)
                nlast = node.right
            pre_node.next = node
            if node == last:
                # 遇到换层时置空前置节点，避免粘连
                pre_node = Node()
                last = nlast
            else:
                pre_node = node
            
        return root
