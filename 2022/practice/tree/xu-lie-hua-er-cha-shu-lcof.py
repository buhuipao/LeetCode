# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
# 按层序列或者反序列
class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        queue = [root]
        ret = []
        while queue:
            node = queue[0]
            queue = queue[1:]
            if not node:
                ret.append("null")
                continue
            ret.append(str(node.val))
            queue.append(node.left)
            queue.append(node.right)

        return ','.join(ret)
        
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data or data == "null":
            return None
        nodes = data.split(",")
        root = TreeNode(int(nodes[0]))
        
        queue, i = [root], 1
        while i < len(nodes):
            parent = queue[0]
            queue = queue[1:]

            left = nodes[i] 
            i += 1
            if left != "null":
                parent.left = TreeNode(int(left))
                queue.append(parent.left)
            
            right = nodes[i]
            i += 1
            if right != "null":
                parent.right = TreeNode(int(right))
                queue.append(parent.right)

        return root
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
