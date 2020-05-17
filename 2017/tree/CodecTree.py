# _*_ coding: utf-8 _*_


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec(object):
    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return []

        _queue = [root]
        result = []
        # 按层遍历二叉树
        while _queue:
            node = _queue.pop(0)
            # 无论是否有左右儿子，都压入队列
            # 如果为空用＃进行占位, 但不压入队列
            if node:
                result.append(node.val)
                _queue.append(node.left)
                _queue.append(node.right)
            else:
                result.append('#')
        return result

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None

        root = TreeNode(data[0])
        _queue = [root]
        node = root
        i = 1
        # 每次进行左右儿子的生成
        while i+1 < len(data):
            node = _queue.pop(0)
            if data[i] != '#':
                node.left = TreeNode(data[i])
                _queue.append(node.left)
            if data[i+1] != '#':
                node.right = TreeNode(data[i+1])
                _queue.append(node.right)
            i += 2
        return root

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
