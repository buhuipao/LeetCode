class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec:

    def serialize1(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        按层遍历
        """
        if not root:
            return ''
        _queue = [root]
        result = []
        while _queue:
            node = _queue.pop(0)
            if node:
                result.append(node.val)
                _queue.append(node.left)
                _queue.append(node.right)
            else:
                result.append('#')
        return '$'.join(map(str, result))

    def deserialize1(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        按层反序列化
        """
        if not data:
            return None
        data = data.split('$')
        root = TreeNode(data[0])
        _queue = [root]
        for i in xrange(1, len(data), 2):
            node = _queue.pop(0)
            if data[i] != '#':
                node.left = TreeNode(data[i])
                _queue.append(node.left)
            if data[i+1] != '#':
                node.right = TreeNode(data[i+1])
                _queue.append(node.right)
        return root

    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        先序遍历
        """
        if not root:
            return ''
        _stack = [root]
        node = root
        result = []
        while _stack:
            node = _stack.pop(-1)
            if node:
                result.append(str(node.val))
                _stack.append(node.right)
                _stack.append(node.left)
            else:
                result.append('#')
        return '$'.join(result)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        先序反序列化
        """
        if not data:
            return
        data = data.split('$')

        root = TreeNode(int(data[0]))
        node = root
        right = False
        _stack = [root]
        for i in xrange(1, len(data)-2, 1):
            if data[i] != '#':
                if not right:
                    node.left = TreeNode(int(data[i]))
                    node = node.left
                else:
                    node.right = TreeNode(int(data[i]))
                    node = node.right
                _stack.append(node)
                right = False
            else:
                # 更新标记点，指导树枝往右生长
                right = True
                node = _stack.pop(-1)
        return root
# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
